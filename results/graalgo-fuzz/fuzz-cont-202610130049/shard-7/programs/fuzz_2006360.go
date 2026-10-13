package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0() string {
    switch "a" {
    case "hello":
        switch (("" + "") + "") {
        case "hello":
            switch "foo bar" {
            case "foo bar":
                var v5 int16 = -961
                _ = v5
            default:
                var v6 bool = (((1000 == 32767) || true) && true)
                _ = v6
            }
        case "":
            if (42 >= -2147483648) {
                var v7 int64 = 42
                v7 -= (v7 % (((v7 / ((v7) | 1))) | 1))
                _ = v7
            } else {
                v8 := []int{127, 9223372036854775807, 32767}
                _ = v8
            }
        }
    }
    var v9 int8 = -1
    for fi10 := 0; fi10 < 3; fi10++ {
        fmt.Println(fi10)
        _ = fi10
    }
    for fi11 := 0; fi11 < 18; fi11++ {
        v12 := []int{42, (fi11 - fi11)}
        _ = fi11
        _ = v12
    }
    _ = v9
    return ""
}

func main() {
    v1 := make(map[string]bool)
    v2 := []string{("a" + "foo bar"), "\n\t"}
    var v3 string = (("" + "") + "\n\t")
    v4 := make(map[int]uint32)
    var v13 int32 = 1
    v3 = ""
    v13 = 1
    _ = len(v2)
    var v14 uint64 = 0
    fmt.Println(v14)
    v15 := GenS1{F0: -333}
    v14 = 759
    v14 = 551
    var v16 uint32 = 0
    fmt.Println(v16, v16)
    fmt.Println(v13)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
