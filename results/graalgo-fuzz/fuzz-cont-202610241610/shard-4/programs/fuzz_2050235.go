package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint; F1 uint64; F2 bool }
type GenS3 struct { F0 int16; F1 string; F2 int32 }
func genFn0() uint64 {
    if true {
        var v7 int16 = 339
        v8 := []int{-794}
        _ = v7
        _ = v8
    } else {
        var v9 int = 0
        if ((("" + "") + "\n\t") != "hello") {
            if ("hello" > ("hello" + "a")) {
                fmt.Println(v9)
            }
            v9 = -144
            v10 := GenS1{F0: 676}
            _ = v10
        }
        _ = v9
    }
    v11 := GenS1{F0: 2}
    for fi12 := 0; fi12 < 2; fi12++ {
        fmt.Println(fi12)
        v13 := GenS1{F0: 2}
        _ = fi12
        _ = v13
    }
    var v14 bool = ((!(-2147483648 >= 1)) && false)
    var v15 uint = 996
    _ = v11
    _ = v14
    _ = v15
    return 1
}

func main() {
    var v1 uint16 = 255
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]int)
    var v4 uint32 = 42
    v5 := make(map[int]string)
    v6 := make(map[uint32]string)
    v16 := GenS2{F0: 42, F1: 452, F2: ((true || true) || false)}
    fmt.Println(v1)
    var v17 float64 = (1e10 * (0.1 / 0.1))
    if true {
        var v18 string = "foo bar"
        _ = v18
    }
    for fi19 := 0; fi19 < 4; fi19++ {
        if (!(-1 == -333)) {
            if (v4 >= v4) {
                var v20 uint = 0
                _ = v20
            }
        }
        v21 := GenS2{F0: 4294967295, F1: 579, F2: (((1e10 - -2.25) <= 1e10) && (("a" + ("" + ("" + "hello"))) >= "\n\t"))}
        _ = fi19
        _ = v21
    }
    v22 := GenS1{F0: (v4 ^ v4)}
    switch "a" {
    case "":
        v23 := []int{2147483647, -280}
        _ = v23
    case "foo bar":
        v24 := GenS3{F0: 42, F1: "hello", F2: -1}
        _ = v24
    }
    var v25 int64 = -1
    if false {
        fmt.Println(v25, v25, v17)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
