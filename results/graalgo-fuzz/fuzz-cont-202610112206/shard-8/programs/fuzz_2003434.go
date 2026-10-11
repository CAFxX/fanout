package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0() uint16 {
    var v5 uint32 = 2
    var v6 int16 = 32767
    v7 := []int{}
    for fi8 := 0; fi8 < 13; fi8++ {
        var v9 int = (^(fi8))
        switch 4294967295 {
        case -263:
            _ = len(v7)
        }
        _ = len(v7)
        _ = fi8
        _ = v9
    }
    var v10 bool = ((("a" + ("a" + "foo bar")) >= (((((("\n\t" + ((("foo bar" + "") + (("" + "") + "foo bar")) + ("foo bar" + "foo bar"))) + "") + "foo bar") + "") + "hello") + (("hello" + "") + "a"))) && true)
    _ = v10
    _ = v5
    _ = v6
    _ = v7
    return 711
}
func genFn1() int64 {
    var v11 int16 = 1000
    v11 -= (^(v11))
    if false {
        if true {
            var v12 int32 = -972
            _ = v12
        } else {
            if false {
                if false {
                    v11 -= (v11 - v11)
                    v13 := GenS1{F0: 2}
                    _ = v13
                }
            } else {
                var v14 uint16 = 91
                v11 = -1
                switch "foo bar" {
                case "a":
                    v15 := []int{}
                    _ = v15
                }
                _ = v14
            }
        }
        v16 := []int{-9223372036854775808}
        _ = v16
    } else {
        var v17 uint64 = 1
        _ = v17
    }
    v11 += v11
    fmt.Println(v11)
    _ = v11
    return 0
}

func main() {
    var v1 int16 = 42
    v2 := make(map[int]string)
    var v3 int8 = 0
    v4 := []int{}
    var v18 uint8 = 128
    v3 -= 127
    for fi19 := 0; fi19 < 5; fi19++ {
        v20 := GenS1{F0: 0}
        if ("a" < "a") {
            var v21 uint32 = 2147483648
            _ = v21
        } else {
            v18 = (v18 + v18)
            v3 += v3
        }
        var v22 uint8 = v18
        _ = fi19
        _ = v20
        _ = v22
    }
    if (int8(v1) == v3) {
        fmt.Println(v18)
    }
    var v23 int = 128
    fmt.Println(v3)
    v23 -= -1
    var v24 uint = 1
    for fi25 := 0; fi25 < 9; fi25++ {
        v3 += (-(v3))
        _ = fi25
    }
    v4 = append(v4, v23)
    v24 = v24
    if true {
        v26 := v2[v23]
        v3 -= int8(v23)
        v3 = v3
        _ = v26
    } else {
        if false {
            v3 += v3
        }
        v27 := GenS1{F0: (v18 ^ 153)}
        v28 := GenS1{F0: 255}
        _ = v27
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
