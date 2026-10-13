package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8 }
func genFn0() int16 {
    for fi6 := 0; fi6 < 20; fi6++ {
        var v7 bool = true
        switch "\n\t" {
        case "foo bar":
            v7 = false
        }
        _ = fi6
        _ = v7
    }
    if ("hello" >= "foo bar") {
        var v8 uint = 455
        fmt.Println(v8)
        var v9 string = ("" + (("\n\t" + "foo bar") + "\n\t"))
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: 0}
    v11 := []int{127, 32767, -2147483648, -91}
    var v12 string = "foo bar"
    _ = v10
    _ = v11
    _ = v12
    return 32767
}
func genFn1(p0 bool, p1 uint32) uint8 {
    var v13 int16 = 32767
    p1 = p1
    p0 = false
    if false {
        for fi14 := 0; fi14 < 3; fi14++ {
            v13 = v13
            _ = fi14
        }
    }
    switch 127 {
    case -1:
        v15 := []int{-9223372036854775808, 1, 32768}
        _ = v15
    }
    _ = p0
    _ = p1
    _ = v13
    return 0
}

func main() {
    v1 := make(map[string]string)
    var v2 string = "foo bar"
    v3 := make(map[int]uint32)
    var v4 int = 802
    var v5 string = v2
    for fi16 := 0; fi16 < 11; fi16++ {
        fmt.Println(v4, fi16, v2)
        _ = fi16
    }
    v2 = ("a" + "foo bar")
    if (42 != 42) {
        for fi17 := 0; fi17 < 10; fi17++ {
            if (!false) {
                v5 = v2
                v4 -= fi17
            }
            for fi18 := 0; fi18 < 5; fi18++ {
                var v19 int8 = 127
                var v20 uint64 = 42
                _ = fi18
                _ = v19
                _ = v20
            }
            fmt.Println(v2)
            _ = fi17
        }
    }
    v2 = ("a" + (v2 + v2))
    var v21 int = (v4 * (^(v4)))
    v4 += v21
    var v22 uint8 = 128
    fmt.Println(v2, v4, v22)
    switch v21 {
    case -19:
        v23 := GenS2{F0: (v22 << ((99) & 63))}
        _ = v23
    case 32767:
        switch v4 {
        case -2147483648:
            var v24 int16 = 1000
            _ = v24
        case -102:
            v25 := []int{v4, 0}
            _ = v25
        }
    case 4294967295:
        switch (v21 & (v21 << ((4294967295) & 63))) {
        case -250:
            v21 -= v21
        case -1:
            v4 = (v4 >> ((18446744073709551615) & 63))
        case 753:
            switch v2 {
            case "":
                if false {
                    v2 = v5
                } else {
                    v2 = (v5 + "a")
                    var v26 int8 = -128
                    _ = v26
                }
            case "\n\t":
                v27 := GenS1{F0: 42}
                _ = v27
            default:
                var v28 string = ("a" + (v5 + "\n\t"))
                _ = v28
            }
        }
    }
    if true {
        for fi29 := 0; fi29 < 13; fi29++ {
            v30 := []int{(v4 - 42), v21}
            fmt.Println(v5, v4)
            _ = fi29
            _ = v30
        }
        v31 := []int{v21, v21}
        if false {
            v32 := GenS2{F0: v22}
            var v33 string = "hello"
            fmt.Println(v2, v33, v5)
            _ = v32
            _ = v33
        }
        _ = v31
    } else {
        v34 := GenS2{F0: v22}
        var v35 int = (v4 ^ v21)
        _ = v34
        _ = v35
    }
    switch v2 {
    case "\n\t":
        var v36 bool = false
        _ = v36
    case "hello":
        v21 = v4
    case "":
        fmt.Println(v4, v22, v2)
    }
    var v37 uint32 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v37
    _ = v4
    _ = v5
}
