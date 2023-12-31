---
date: "2018-06-01T19:00:00+02:00"
title: "合并请求"
slug: "pull-request"
sidebar_position: 13
toc: false
draft: false
aliases:
  - /zh-cn/pull-request
menu:
  sidebar:
    parent: "usage"
    name: "Pull Request"
    sidebar_position: 13
    identifier: "pull-request"
---

# 合并请求

合并请求(PR)是一种提出对仓库进行更改的方式。
它是一种将一个分支合并到另一个分支的请求，附带有对所做更改的描述。
合并请求通常用作贡献者对仓库贡献代码的方式，仓库的维护者可以通过对合并请求进行审查来决定是否接受这些更改。

## 创建合并请求

要创建合并请求，您需要遵循以下步骤：

1. **Fork 仓库** - 如果您没有直接对仓库进行更改的权限，您需要将仓库 fork 到您自己的账户中。
这将创建一个您可以对其进行更改的仓库副本。

2. **创建分支（可选）** - 在 fork 的仓库中创建一个新分支，该分支包含您要提出的更改。
给分支取一个描述性的名称，以指示更改的内容。

3. **进行更改** - 进行您想要的更改，提交并将其推送到 fork 的仓库中。

4. **创建合并请求** - 转到原始仓库并转到“合并请求”选项卡。单击“新建合并请求”按钮，并将您的新分支选择为源分支。
为您的合并请求输入描述性标题和描述，然后单击“创建合并请求”。

## 评审合并请求

创建合并请求后，将触发评审流程。仓库的维护者将收到合并请求的通知，并可以审查所做的更改。
他们可以留下评论、请求更改或批准更改。

如果维护者请求更改，您需要在分支中进行这些更改，并将更改推送到 fork 的仓库中。
合并请求将自动使用新更改进行更新。

如果维护者批准更改，他们可以将合并请求合并到仓库中。

## 关闭合并请求

如果您不接受该合并请求，您可以关闭它。
要关闭合并请求，请转到打开的合并请求并单击“关闭合并请求”按钮。这将关闭合并请求并且不会将其合并。

## 使用“Work In Progress”标记

在合并请求中使用“Work In Progress”标记可以防止合并请求被意外合并。
要将合并请求标记为“Work In Progress”，您必须在其标题中添加前缀`WIP:`或`[WIP]`（不区分大小写）。
标记前缀可以在您的`app.ini`文件中进行配置：

```
[repository.pull-request]
WORK_IN_PROGRESS_PREFIXES=WIP:,[WIP]
```

列表的第一个值将用于 helpers 程序。

## 合并请求模板

有关合并请求模板的更多信息请您移步 : [工单与合并请求模板](issue-pull-request-templates)
