package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint32 }
func genFn0(p0 uint64) uint32 {
    if (((0.0 - 1e10) - float64(float32(0.0))) < 1e10) {
        var v5 int8 = 0
        p0 += (p0 << ((682) & 63))
        _ = v5
    } else {
        if ((true && (-9223372036854775808 == 1)) || (1 <= 0)) {
            if ("hello" <= "a") {
                if (-1 >= 1) {
                    var v6 int32 = -2147483648
                    var v7 int8 = 42
                    _ = v6
                    _ = v7
                } else {
                    v8 := GenS1{F0: ((("foo bar" + ((("a" + "") + "") + ("a" + ("a" + "foo bar")))) + "foo bar") + ("foo bar" + "")), F1: -819}
                    p0 += p0
                    _ = v8
                }
                for fi9 := 0; fi9 < 8; fi9++ {
                    v10 := []int{fi9, fi9}
                    _ = fi9
                    _ = v10
                }
                var v11 int = -1
                _ = v11
            }
            v12 := GenS1{F0: ("foo bar" + "a"), F1: 42}
            _ = v12
        }
    }
    v13 := GenS1{F0: "a", F1: 0}
    v14 := GenS2{F0: 1000}
    _ = p0
    _ = v13
    _ = v14
    return 42
}

func main() {
    var v1 bool = (-2147483648 >= 197)
    v2 := []uint32{1, 1, 42, 4294967295}
    v3 := make(map[int]string)
    var v4 int64 = 844
    var v15 string = "foo bar"
    for fi16 := 0; fi16 < 15; fi16++ {
        v4 -= -823
        fmt.Println(fi16, fi16)
        _ = fi16
    }
    for fi17 := 0; fi17 < 8; fi17++ {
        var v18 int8 = 127
        var v19 float64 = float64(1.5)
        v20 := GenS3{F0: 742}
        _ = fi17
        _ = v18
        _ = v19
        _ = v20
    }
    for fi21 := 0; fi21 < 13; fi21++ {
        switch ("" + ("" + v15)) {
        case "a":
            for fi22 := 0; fi22 < 17; fi22++ {
                v23 := GenS4{F0: genFn0(4294967295)}
                v24 := GenS3{F0: 2}
                _ = fi22
                _ = v23
                _ = v24
            }
        case "":
            var v25 int8 = 0
            _ = v25
        case "\n\t":
            var v26 uint32 = 2147483648
            _ = v26
        }
        var v27 float64 = 3.14159
        _ = len(v3)
        _ = fi21
        _ = v27
    }
    v1 = v1
    fmt.Println(v15, v15, v15)
    v4 = v4
    if true {
        v28 := GenS1{F0: v15, F1: -1}
        _ = v28
    } else {
        v15 = v15
    }
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
