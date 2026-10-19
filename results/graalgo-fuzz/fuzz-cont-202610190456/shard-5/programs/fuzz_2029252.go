package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 bool }
func genFn0(p0 int8, p1 bool, p2 int32) uint32 {
    p0 = p0
    p0 += -1
    v4 := GenS1{F0: 42, F1: 779}
    switch (("a" + ("a" + ("hello" + "a"))) + "a") {
    case "hello":
        if p1 {
            v5 := []int{897, -2147483648, 224}
            var v6 float64 = 0.1
            _ = v5
            _ = v6
        }
    case "\n\t":
        if false {
            var v7 string = "a"
            _ = v7
        }
    case "a":
        for fi8 := 0; fi8 < 2; fi8++ {
            fmt.Println(fi8)
            p1 = p1
            _ = fi8
        }
    }
    fmt.Println(p0, p2, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 26
}
func genFn1(p0 int32, p1 uint16) int64 {
    switch 9223372036854775807 {
    case 2147483647:
        var v9 uint16 = (p1 & (^(p1)))
        _ = v9
    case 4294967295:
        v10 := []int{42, -1}
        _ = v10
    }
    var v11 uint16 = (p1 | p1)
    v12 := GenS1{F0: 386, F1: 94}
    switch "foo bar" {
    case "a":
        v13 := []int{384, 42, 481, 323}
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return 9223372036854775807
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint8 = 76
    var v3 int16 = 74
    for fi14 := 0; fi14 < 10; fi14++ {
        delete(v1, 4294967295)
        switch "\n\t" {
        case "\n\t":
            v3 = v3
        case "hello":
            fmt.Println(v2, v3, v2)
        case "":
            var v15 int8 = 100
            _ = v15
        default:
            for fi16 := 0; fi16 < 20; fi16++ {
                fmt.Println(v2, fi14, fi14)
                v17 := []int{int(v3), fi14, (-(fi14))}
                _ = fi16
                _ = v17
            }
        }
        v2 += (^(v2))
        _ = fi14
    }
    v18 := v1[2147483647]
    for fi19 := 0; fi19 < 15; fi19++ {
        v20 := []int{}
        for fi21 := 0; fi21 < 18; fi21++ {
            var v22 uint = 255
            var v23 int64 = -996
            switch -2147483648 {
            case 127:
                var v24 int32 = -359
                _ = v24
            case 42:
                _ = len(v1)
            }
            switch "\n\t" {
            case "hello":
                if v18 {
                    switch (((("hello" + (("" + "hello") + (("foo bar" + ("\n\t" + "foo bar")) + ("hello" + "a")))) + "a") + ("foo bar" + "foo bar")) + "hello") {
                    case "hello":
                        var v25 float64 = 3.14159
                        _ = v25
                    }
                    if (int16(fi21) != v3) {
                        if len(v20) >= 2 {
                            v20 = v20[1:len(v20)]
                        }
                    }
                    var v26 bool = true
                    _ = v26
                } else {
                    v27 := GenS1{F0: 42, F1: 1000}
                    _ = v27
                }
            case "":
                var v28 uint64 = 331
                _ = v28
            case "\n\t":
                v18 = v18
            default:
                v29 := GenS1{F0: 529, F1: v3}
                _ = v29
            }
            _ = fi21
            _ = v22
            _ = v23
        }
        _ = fi19
        _ = v20
    }
    v30 := GenS2{F0: v18}
    if (!v18) {
        v18 = (v18 && false)
    } else {
        fmt.Println(v3, v18, v3)
        var v31 int8 = 11
        _ = v31
    }
    switch "foo bar" {
    case "hello":
        var v32 float32 = 0.1
        _ = v32
    case "a":
        var v33 int64 = 9223372036854775807
        _ = v33
    case "":
        if false {
            fmt.Println(v18, v18)
            var v34 int16 = 469
            _ = v34
        }
    }
    switch "\n\t" {
    case "hello":
        fmt.Println(v3, v2, v3)
    case "a":
        v35 := v1[genFn0(101, v18, 133)]
        _ = v35
    case "":
        fmt.Println(v18, v3)
    }
    fmt.Println(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v30
}
