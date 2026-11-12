package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 uint64, p1 uint8) int32 {
    v6 := []int{-1}
    p1 = (^(p1))
    v7 := GenS1{F0: 1}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return -2147483648
}

func main() {
    var v1 uint16 = 42
    var v2 uint16 = v1
    var v3 uint8 = 2
    v4 := make(map[uint32]int)
    var v5 int16 = -32768
    switch 418 {
    case 127:
        for fi8 := 0; fi8 < 18; fi8++ {
            fmt.Println(v2)
            var v9 int = (-(fi8))
            _ = fi8
            _ = v9
        }
    }
    v10 := v4[599]
    if (true && (2147483648 > 459)) {
        v11 := GenS1{F0: 0}
        var v12 int8 = -128
        var v13 int64 = 0
        _ = v11
        _ = v12
        _ = v13
    }
    switch (((("foo bar" + (("foo bar" + ("hello" + ("hello" + "\n\t"))) + "")) + "foo bar") + "\n\t") + "a") {
    case "foo bar":
        v14 := []int{}
        _ = v14
    case "hello":
        delete(v4, 2147483648)
    case "":
        if true {
            v3 -= 234
            var v15 uint64 = 2
            v1 -= uint16(v15)
            _ = v15
        }
    }
    if ((v10 / (((^(v10))) | 1)) != (-(v10))) {
        for fi16 := 0; fi16 < 13; fi16++ {
            var v17 float32 = (1.5 / 1.0)
            _ = fi16
            _ = v17
        }
        var v18 int8 = -128
        if false {
            v18 = (v18 >> ((1) & 63))
            v19 := []int{4294967295, (v10 & 128)}
            _ = v19
        }
        _ = v18
    } else {
        v20 := GenS1{F0: 0}
        _ = v20
    }
    fmt.Println(v5, v10)
    v5 -= 0
    var v21 bool = true
    switch (("foo bar" + ((((("\n\t" + ("" + "\n\t")) + (("\n\t" + ("hello" + "a")) + "\n\t")) + "\n\t") + ("" + "a")) + (("" + ("a" + ("\n\t" + "hello"))) + ("hello" + ("" + "hello"))))) + "\n\t") {
    case "":
        v1 += 202
    }
    for fi22 := 0; fi22 < 19; fi22++ {
        v23 := GenS1{F0: 105}
        _ = fi22
        _ = v23
    }
    var v24 int32 = 0
    var v25 int16 = v5
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
