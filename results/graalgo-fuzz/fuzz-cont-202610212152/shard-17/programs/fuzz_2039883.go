package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint; F1 int8 }
func genFn0(p0 int8, p1 uint16, p2 int16) uint8 {
    if (569 > 642) {
        switch (("\n\t" + (("a" + (("\n\t" + "") + "\n\t")) + "\n\t")) + (("\n\t" + ("" + "a")) + ((((("foo bar" + "a") + ("foo bar" + "foo bar")) + ("" + "hello")) + ("hello" + "hello")) + "foo bar"))) {
        case "foo bar":
            p2 = (p2 + (^(p2)))
        }
        var v4 uint32 = 1
        _ = v4
    }
    if true {
        for fi5 := 0; fi5 < 9; fi5++ {
            for fi6 := 0; fi6 < 12; fi6++ {
                var v7 uint = 450
                _ = fi6
                _ = v7
            }
            if true {
                p1 = 65535
                var v8 uint = 255
                _ = v8
            } else {
                v9 := GenS1{F0: 128, F1: 4294967295}
                v10 := []int{}
                _ = v10
                _ = v9
            }
            v11 := []int{0, -568, fi5}
            _ = fi5
            _ = v11
        }
    }
    fmt.Println(p2, p0, p2)
    var v12 int = 32768
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 255
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []int{2147483647, -873}
    v3 := make(map[string]bool)
    delete(v3, "foo bar")
    var v13 uint16 = 255
    v14 := v3["a"]
    delete(v1, "")
    switch 42 {
    case 936:
        var v15 int = -2147483648
        _ = v15
    case 1:
        v13 = (v13 % (((^(v13))) | 1))
    case -1:
        var v16 uint64 = 1
        _ = v16
    }
    v14 = false
    v17 := GenS2{F0: 2147483647}
    var v18 bool = v14
    for fi19 := 0; fi19 < 5; fi19++ {
        v13 -= 1
        var v20 uint = 1
        _ = fi19
        _ = v20
    }
    for fi21 := 0; fi21 < 14; fi21++ {
        for fi22 := 0; fi22 < 12; fi22++ {
            var v23 uint = 255
            v13 -= (v13 & (v13 & v13))
            _ = fi22
            _ = v23
        }
        _ = fi21
    }
    if (-2147483648 == -9223372036854775808) {
        var v24 int32 = 42
        if ((0.0 >= 3.14159) || v14) {
            if len(v2) > 0 {
                v2[0] = 32768
            }
        } else {
            v25 := GenS2{F0: 4294967295}
            v14 = (!(0.1 > 1e10))
            _ = v25
        }
        _ = v24
    } else {
        if false {
            v13 = v13
            v1["a"] = 106
        } else {
            if (!v14) {
                var v26 uint16 = v13
                _ = v26
            } else {
                v13 -= v13
            }
            v13 += 1
            v27 := GenS1{F0: 1, F1: 4294967295}
            _ = v27
        }
    }
    v13 -= (v13 % (((v13 * v13)) | 1))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi28 := 0; fi28 < 7; fi28++ {
        v29 := GenS3{F0: 18446744073709551615, F1: -128}
        v18 = v18
        _ = fi28
        _ = v29
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
