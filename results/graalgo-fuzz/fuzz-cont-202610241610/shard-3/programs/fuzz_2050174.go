package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32; F1 bool; F2 bool }
type GenS3 struct { F0 int; F1 uint }
func genFn0(p0 string) int {
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 uint8 = 127
        _ = fi7
        _ = v8
    }
    fmt.Println(p0)
    for fi9 := 0; fi9 < 17; fi9++ {
        var v10 int64 = -9223372036854775808
        v11 := GenS1{F0: 42}
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 uint8 = 228
    for fi13 := 0; fi13 < 6; fi13++ {
        v14 := GenS2{F0: -360, F1: true, F2: false}
        _ = fi13
        _ = v14
    }
    _ = p0
    _ = v12
    return 128
}
func genFn1(p0 int64, p1 uint8, p2 int8) uint32 {
    for fi15 := 0; fi15 < 11; fi15++ {
        for fi16 := 0; fi16 < 12; fi16++ {
            if true {
                p0 = -303
                var v17 string = "\n\t"
                _ = v17
            } else {
                var v18 float32 = 1e10
                v19 := []int{}
                _ = v18
                _ = v19
            }
            var v20 uint64 = 0
            v21 := []int{genFn0(((("hello" + "hello") + "a") + (("hello" + "\n\t") + ("hello" + ((("foo bar" + "foo bar") + "a") + ((("a" + "foo bar") + (("" + ("\n\t" + "hello")) + (("foo bar" + ("" + "hello")) + ""))) + "hello"))))))}
            _ = fi16
            _ = v20
            _ = v21
        }
        _ = fi15
    }
    if true {
        var v22 int32 = -2147483648
        _ = v22
    }
    var v23 uint32 = 708
    _ = p0
    _ = p1
    _ = p2
    _ = v23
    return (^(v23))
}

func main() {
    var v1 bool = (255 != 65535)
    var v2 string = "a"
    var v3 int8 = 121
    v4 := make(map[uint32]int)
    v5 := make(map[int]int)
    v6 := make(map[uint32]bool)
    switch v2 {
    case "":
        v24 := v4[4294967295]
        _ = v24
    default:
        fmt.Println(v3, v1)
    }
    var v25 uint32 = 1
    var v26 uint16 = 42
    if v1 {
        fmt.Println(v3, v25)
        var v27 uint8 = 0
        for fi28 := 0; fi28 < 16; fi28++ {
            if false {
                fmt.Println(v25, v2, v3)
            } else {
                fmt.Println(v26)
            }
            v3 = v3
            v2 = (("\n\t" + v2) + "")
            _ = fi28
        }
        _ = v27
    }
    v3 -= 0
    var v29 int32 = 0
    v6[(v25 % ((2147483647) | 1))] = (!(genFn1(1, 248, 42) > (v25 << ((630) & 63))))
    v30 := []int{865}
    var v31 uint = 255
    var v32 uint64 = 325
    delete(v4, 0)
    for fi33 := 0; fi33 < 19; fi33++ {
        v34 := GenS3{F0: genFn0("hello"), F1: v31}
        var v35 bool = v1
        _ = fi33
        _ = v34
        _ = v35
    }
    for fi36 := 0; fi36 < 14; fi36++ {
        for fi37 := 0; fi37 < 15; fi37++ {
            var v38 bool = v1
            var v39 int64 = -1
            switch (((v2 + (v2 + v2)) + v2) + (v2 + v2)) {
            case "a":
                if ((v31 * 0) <= (^(v31))) {
                    var v40 float64 = -2.25
                    _ = v40
                }
            case "\n\t":
                v41 := GenS2{F0: v29, F1: true, F2: false}
                _ = v41
            case "foo bar":
                v29 -= v29
            }
            v42 := v5[(fi36 | int(v31))]
            _ = fi37
            _ = v38
            _ = v39
            _ = v42
        }
        var v43 string = "foo bar"
        _ = fi36
        _ = v43
    }
    v44 := GenS2{F0: (^(v29)), F1: (!((115 != 255) && (76 >= 42))), F2: ((v29 >> ((v31) & 63)) <= 1)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v44
    _ = v5
    _ = v6
}
