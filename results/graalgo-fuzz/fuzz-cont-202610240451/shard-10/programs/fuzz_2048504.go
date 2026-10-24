package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 int64; F1 uint; F2 uint64 }
func genFn0(p0 int16) uint64 {
    var v5 int16 = (p0 * (^(p0)))
    v6 := GenS1{F0: -1, F1: (("\n\t" + (("a" + (("\n\t" + "\n\t") + (("a" + ("" + (("\n\t" + "\n\t") + "foo bar"))) + "a"))) + "foo bar")) + "a")}
    _ = p0
    _ = v5
    _ = v6
    return 932
}
func genFn1() int32 {
    if (42 != 506) {
        var v7 uint16 = 255
        switch "\n\t" {
        case "":
            var v8 int8 = 42
            _ = v8
        }
        _ = v7
    } else {
        var v9 int32 = 1
        fmt.Println(v9)
        v9 = v9
        _ = v9
    }
    if (105 == 1) {
        if (127 > 2147483647) {
            v10 := GenS1{F0: 0, F1: ("" + "a")}
            var v11 string = "hello"
            _ = v10
            _ = v11
        }
    }
    var v12 uint64 = genFn0(662)
    fmt.Println(v12)
    if false {
        if (32767 >= 65535) {
            v12 += (^(v12))
            v12 -= 4294967295
            var v13 string = ""
            _ = v13
        } else {
            v14 := GenS2{F0: 1, F1: 1, F2: v12}
            _ = v14
        }
        v15 := GenS1{F0: -9223372036854775808, F1: "hello"}
        v16 := []int{42, -2147483648, -9223372036854775808, 65}
        _ = v15
        _ = v16
    } else {
        switch "a" {
        case "a":
            v12 -= (v12 ^ 0)
        case "":
            switch -9223372036854775808 {
            case 0:
                for fi17 := 0; fi17 < 5; fi17++ {
                    v18 := []int{fi17, fi17, 1, (fi17 / ((fi17) | 1))}
                    _ = fi17
                    _ = v18
                }
            }
        default:
            fmt.Println(v12)
        }
        switch ("\n\t" + "") {
        case "":
            var v19 uint8 = 127
            _ = v19
        case "a":
            if (0 < 1) {
                v20 := []int{int(v12), 902, -2147483648, 2147483647}
                _ = v20
            } else {
                for fi21 := 0; fi21 < 4; fi21++ {
                    v12 += v12
                    _ = fi21
                }
                v22 := GenS2{F0: 0, F1: 86, F2: v12}
                v23 := GenS1{F0: 83, F1: ((("" + ("\n\t" + ("" + "a"))) + "\n\t") + ("hello" + "hello"))}
                _ = v22
                _ = v23
            }
        default:
            var v24 uint32 = 42
            _ = v24
        }
    }
    _ = v12
    return 42
}

func main() {
    var v1 uint = 2
    v2 := []uint32{2}
    var v3 uint8 = 127
    var v4 int32 = 42
    v25 := GenS2{F0: 0, F1: v1, F2: 0}
    if (3.14159 >= 0.1) {
        var v26 uint8 = (v3 - 104)
        v1 += (^(v1))
        _ = v26
    }
    v1 = v1
    switch 32768 {
    case 0:
        v4 -= (v4 / ((genFn1()) | 1))
    }
    switch (("a" + "a") + "hello") {
    case "":
        var v27 int64 = 9223372036854775807
        _ = v27
    case "hello":
        v4 -= genFn1()
    case "foo bar":
        v1 = v1
    }
    if (127 <= 42) {
        _ = len(v2)
    }
    var v28 string = (("" + "") + (("" + "") + "\n\t"))
    fmt.Println(v4)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v28
    _ = v3
    _ = v4
}
