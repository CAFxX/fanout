package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 string; F2 uint64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int; F1 uint16; F2 int8 }
func genFn0(p0 string) int {
    v4 := GenS1{F0: 1}
    p0 = p0
    p0 = p0
    _ = p0
    _ = v4
    return 746
}
func genFn1(p0 bool, p1 int16) uint64 {
    p0 = p0
    if false {
        switch "foo bar" {
        case "a":
            var v5 int32 = -966
            _ = v5
        case "foo bar":
            for fi6 := 0; fi6 < 14; fi6++ {
                fmt.Println(p0)
                var v7 uint = 65535
                _ = fi6
                _ = v7
            }
        case "\n\t":
            v8 := []int{2147483647, genFn0("foo bar")}
            _ = v8
        }
        switch "a" {
        case "":
            v9 := []int{42, genFn0(("a" + "\n\t")), 417}
            _ = v9
        case "hello":
            for fi10 := 0; fi10 < 3; fi10++ {
                p1 += (^(p1))
                var v11 int8 = 0
                v12 := []int{(fi10 ^ (fi10 / (((fi10 % (((fi10 ^ fi10)) | 1))) | 1))), (fi10 + (-(fi10))), (^(fi10)), fi10}
                _ = fi10
                _ = v11
                _ = v12
            }
        case "a":
            p1 = 1000
        default:
            if false {
                var v13 uint16 = 65535
                _ = v13
            } else {
                fmt.Println(p0, p1)
                fmt.Println(p1)
                var v14 float64 = (3.14159 - ((3.14159 / (0.1 * -2.25)) - 0.0))
                _ = v14
            }
        }
        switch "a" {
        case "hello":
            var v15 int32 = 443
            _ = v15
        }
    } else {
        v16 := GenS2{F0: 255, F1: "foo bar", F2: 872}
        if false {
            v17 := GenS3{F0: 5}
            switch -144 {
            case 0:
                for fi18 := 0; fi18 < 14; fi18++ {
                    v19 := []int{(fi18 + genFn0(("a" + ("" + (("hello" + "hello") + ("foo bar" + ("" + ("hello" + ("" + "foo bar"))))))))), (fi18 % ((-9223372036854775808) | 1)), -401, fi18}
                    _ = fi18
                    _ = v19
                }
            }
            _ = v17
        } else {
            switch "foo bar" {
            case "foo bar":
                var v20 int16 = (p1 >> ((2) & 63))
                _ = v20
            case "\n\t":
                v21 := GenS2{F0: 65535, F1: ("hello" + ("\n\t" + (("a" + "a") + ("\n\t" + "hello")))), F2: 685}
                _ = v21
            case "":
                fmt.Println(p1, p0)
            }
            var v22 uint16 = 65535
            _ = v22
        }
        v23 := GenS4{F0: 32767, F1: 65535, F2: 42}
        _ = v16
        _ = v23
    }
    v24 := []int{-158, 1}
    p1 -= p1
    _ = p0
    _ = p1
    _ = v24
    return 441
}

func main() {
    v1 := []uint32{}
    var v2 int64 = -527
    v3 := make(map[int]uint32)
    switch 32767 {
    case 127:
        fmt.Println(v2)
    case -1:
        var v25 int32 = -975
        _ = v25
    case 197:
        fmt.Println(v2)
    default:
        var v26 uint8 = 1
        _ = v26
    }
    for fi27 := 0; fi27 < 6; fi27++ {
        _ = len(v3)
        _ = fi27
    }
    var v28 uint16 = 65535
    v29 := GenS4{F0: 42, F1: v28, F2: -1}
    _ = len(v3)
    v28 -= v28
    for fi30 := 0; fi30 < 6; fi30++ {
        v28 = v28
        switch 9223372036854775807 {
        case 4294967295:
            if (32767 < -1) {
                var v31 int32 = -175
                switch ("\n\t" + ("hello" + "")) {
                case "\n\t":
                    v28 = (v28 & 2)
                case "foo bar":
                    var v32 int = fi30
                    _ = v32
                default:
                    v28 = 133
                }
                _ = v31
            }
        case -1:
            var v33 bool = (!(v28 == v28))
            _ = v33
        default:
            var v34 uint64 = 710
            _ = v34
        }
        _ = fi30
    }
    var v35 uint8 = 2
    var v36 uint8 = v35
    switch 32768 {
    case 0:
        v28 = 1
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v35
    _ = v36
}
