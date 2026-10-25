package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 int32 }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 uint32; F1 bool; F2 uint }
func genFn0() uint {
    var v4 uint16 = 77
    v5 := GenS1{F0: uint64(v4), F1: 4294967295, F2: 2147483647}
    if (false && (((true && false) || false) && false)) {
        fmt.Println(v4)
        var v6 bool = (100 > 42)
        v4 += (v4 | v4)
        _ = v6
    }
    v7 := []int{}
    _ = v4
    _ = v5
    _ = v7
    return 255
}
func genFn1(p0 int16) uint16 {
    for fi8 := 0; fi8 < 16; fi8++ {
        var v9 float32 = (0.0 / (3.14159 * -2.25))
        _ = fi8
        _ = v9
    }
    p0 -= (p0 + p0)
    switch 42 {
    case 42:
        if false {
            p0 += (p0 ^ (p0 ^ (p0 >> ((206) & 63))))
        } else {
            p0 = (p0 >> ((18446744073709551615) & 63))
            if false {
                var v10 int8 = 127
                for fi11 := 0; fi11 < 10; fi11++ {
                    v10 = v10
                    _ = fi11
                }
                v12 := GenS1{F0: 307, F1: -363, F2: 0}
                _ = v10
                _ = v12
            } else {
                for fi13 := 0; fi13 < 3; fi13++ {
                    v14 := []int{fi13}
                    p0 = -750
                    _ = fi13
                    _ = v14
                }
                for fi15 := 0; fi15 < 19; fi15++ {
                    for fi16 := 0; fi16 < 11; fi16++ {
                        var v17 uint16 = 42
                        var v18 uint8 = 255
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                    var v19 uint = 677
                    _ = fi15
                    _ = v19
                }
            }
            p0 = (p0 / ((p0) | 1))
        }
    case 128:
        fmt.Println(p0)
    }
    var v20 int16 = (^(p0))
    for fi21 := 0; fi21 < 15; fi21++ {
        v22 := GenS1{F0: 144, F1: fi21, F2: 1}
        v23 := []int{42, -1, fi21, fi21}
        _ = fi21
        _ = v22
        _ = v23
    }
    _ = p0
    _ = v20
    return 170
}

func main() {
    v1 := make(map[int]string)
    var v2 string = "hello"
    var v3 uint = 4294967295
    if false {
        switch (v2 + v2) {
        case "":
            v3 -= v3
        case "hello":
            v24 := GenS2{F0: 1, F1: (("" + ("a" + v2)) + "a")}
            _ = v24
        }
        v25 := GenS1{F0: 221, F1: 2147483647, F2: 512}
        v26 := GenS2{F0: 244, F1: v2}
        _ = v25
        _ = v26
    } else {
        v27 := []int{}
        _ = v27
    }
    var v28 int32 = 2147483647
    if false {
        var v29 uint64 = 4294967295
        _ = v29
    } else {
        v30 := GenS1{F0: 586, F1: 2147483647, F2: -2147483648}
        v28 += v28
        _ = v30
    }
    if (362 != 32767) {
        fmt.Println(v3, v3, v2)
    }
    if (-1 == -97) {
        v2 = v2
        for fi31 := 0; fi31 < 10; fi31++ {
            var v32 int8 = -119
            _ = fi31
            _ = v32
        }
        v2 = ("a" + ("\n\t" + v2))
    } else {
        fmt.Println(v3)
    }
    v3 -= v3
    v2 = v2
    if (true && (1 < 0)) {
        for fi33 := 0; fi33 < 15; fi33++ {
            switch v2 {
            case "\n\t":
                v28 -= v28
            }
            var v34 int32 = (^(v28))
            _ = fi33
            _ = v34
        }
        if false {
            for fi35 := 0; fi35 < 7; fi35++ {
                fmt.Println(v28, v3)
                _ = fi35
            }
            _ = len(v1)
            v2 = "foo bar"
        }
    } else {
        var v36 int64 = 271
        for fi37 := 0; fi37 < 13; fi37++ {
            v38 := GenS3{F0: 2147483648, F1: true, F2: v3}
            v39 := []int{9223372036854775807, fi37}
            v28 += (v28 & v28)
            _ = fi37
            _ = v38
            _ = v39
        }
        v36 += (v36 - 9223372036854775807)
        _ = v36
    }
    if (false || ((0 < v28) || false)) {
        v40 := []int{786, 127}
        _ = v40
    } else {
        v41 := v1[10]
        _ = v41
    }
    delete(v1, -912)
    var v42 uint8 = 2
    switch ("foo bar" + (v2 + v2)) {
    case "a":
        v43 := v1[99]
        _ = v43
    }
    v2 = v2
    var v44 int8 = 14
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v42
    _ = v44
}
