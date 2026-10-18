package main

import (
    "fmt"
)

func genFn0() bool {
    var v6 bool = false
    var v7 bool = (-9223372036854775808 < -982)
    var v8 int16 = 614
    _ = v6
    _ = v7
    _ = v8
    return (!v7)
}

func main() {
    var v1 float32 = 1.5
    var v2 int8 = 0
    v3 := make(map[string]int)
    var v4 int = 0
    var v5 int32 = -295
    var v9 bool = true
    v9 = ((!(-1 != -1)) || false)
    var v10 string = ((("foo bar" + "\n\t") + "foo bar") + ("hello" + ""))
    v10 = (v10 + "")
    v4 -= 32767
    for fi11 := 0; fi11 < 5; fi11++ {
        v9 = true
        v12 := []int{}
        _ = fi11
        _ = v12
    }
    v1 = (v1 / v1)
    v13 := []int{2147483647}
    for fi14 := 0; fi14 < 18; fi14++ {
        var v15 uint = 374
        _ = fi14
        _ = v15
    }
    v13 = append(v13, 127)
    v2 = (v2 + v2)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
