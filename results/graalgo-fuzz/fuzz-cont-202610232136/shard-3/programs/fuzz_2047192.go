package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0() uint16 {
    v6 := GenS1{F0: true}
    var v7 uint32 = 42
    for fi8 := 0; fi8 < 11; fi8++ {
        v9 := GenS1{F0: (true && (55 < 42))}
        v10 := []int{(-(fi8)), 1}
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = v6
    _ = v7
    return 2
}

func main() {
    v1 := []string{("" + "\n\t"), "\n\t", "hello", ("" + "foo bar")}
    v2 := []int{771}
    v3 := []string{"foo bar", ("hello" + (((("foo bar" + "") + "\n\t") + "\n\t") + ("hello" + ""))), "hello", ("a" + "")}
    var v4 uint16 = 2
    v5 := make(map[uint32]int)
    v4 += 42
    v4 += genFn0()
    v4 -= (v4 & 42)
    fmt.Println(v4)
    var v11 uint16 = v4
    v12 := GenS1{F0: false}
    v11 += (v4 & v11)
    switch -1 {
    case 127:
        var v13 uint8 = 1
        _ = v13
    case 32767:
        var v14 string = ""
        _ = v14
    case 2147483647:
        v15 := GenS1{F0: (2147483647 <= 32767)}
        _ = v15
    }
    v4 += v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
