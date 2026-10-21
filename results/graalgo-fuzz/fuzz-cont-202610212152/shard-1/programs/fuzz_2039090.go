package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int }
func genFn0(p0 int64, p1 int32) uint {
    for fi5 := 0; fi5 < 18; fi5++ {
        fmt.Println(fi5, p1, p0)
        _ = fi5
    }
    for fi6 := 0; fi6 < 10; fi6++ {
        var v7 uint8 = 1
        v8 := []int{fi6, fi6}
        _ = fi6
        _ = v7
        _ = v8
    }
    if (float32(1e10) >= (0.0 - 1.5)) {
        v9 := []int{-596, 127}
        switch "hello" {
        case "":
            for fi10 := 0; fi10 < 5; fi10++ {
                switch 2147483647 {
                case 128:
                    var v11 float32 = -2.25
                    _ = v11
                case 4294967295:
                    var v12 uint64 = 0
                    _ = v12
                case -1:
                    fmt.Println(fi10)
                }
                v13 := GenS1{F0: 42, F1: -75, F2: 2}
                var v14 int64 = p0
                _ = fi10
                _ = v13
                _ = v14
            }
        default:
            var v15 int16 = 32767
            _ = v15
        }
        v16 := GenS1{F0: 290, F1: 1, F2: 42}
        _ = v16
        _ = v9
    } else {
        fmt.Println(p1, p0)
    }
    switch "a" {
    case "":
        p1 = p1
    case "foo bar":
        var v17 float64 = ((0.1 * ((-2.25 / (float64(1.5) * (((1e10 - (3.14159 / 1.5)) + 0.1) / 1.0))) - -2.25)) / 3.14159)
        _ = v17
    default:
        if (-2.25 >= 1.5) {
            fmt.Println(p0, p0)
            v18 := []int{}
            fmt.Println(p0, p1)
            _ = v18
        }
    }
    _ = p0
    _ = p1
    return 4294967295
}
func genFn1(p0 int8, p1 bool) int64 {
    v19 := []int{}
    if p1 {
        var v20 uint64 = 1
        if p1 {
            fmt.Println(p0, v20, v20)
        }
        var v21 uint32 = 65535
        _ = v20
        _ = v21
    } else {
        p0 += (p0 % ((p0) | 1))
        if p1 {
            for fi22 := 0; fi22 < 5; fi22++ {
                fmt.Println(p0, p0, p0)
                _ = fi22
            }
            for fi23 := 0; fi23 < 9; fi23++ {
                p0 = p0
                v24 := GenS2{F0: ("\n\t" + "foo bar")}
                var v25 bool = p1
                _ = fi23
                _ = v24
                _ = v25
            }
            var v26 float32 = (0.1 + ((1e10 + (3.14159 + 0.1)) * ((3.14159 / -2.25) - (0.0 / 1e10))))
            _ = v26
        }
    }
    p1 = false
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v19
    return 948
}

func main() {
    var v1 int = 152
    var v2 int16 = 1000
    var v3 uint = 2
    var v4 float32 = 1.5
    v3 += (v3 * (v3 << ((v3) & 63)))
    v27 := GenS2{F0: (("foo bar" + (("hello" + "a") + "")) + "hello")}
    v28 := GenS3{F0: (v1 & v1)}
    var v29 int = (v1 & v1)
    fmt.Println(v2)
    if ((1e10 / v4) >= v4) {
        var v30 uint = 18446744073709551615
        _ = v30
    }
    for fi31 := 0; fi31 < 13; fi31++ {
        var v32 float64 = (((((0.0 / -2.25) * 0.1) - (-2.25 - ((-2.25 * -2.25) + (-2.25 + 3.14159)))) / float64(1e10)) * -2.25)
        v32 = (3.14159 / 0.1)
        var v33 string = "a"
        _ = fi31
        _ = v32
        _ = v33
    }
    switch (^(v29)) {
    case -962:
        if (((!("foo bar" > (((("\n\t" + ("a" + "a")) + "a") + "") + "\n\t"))) && true) && false) {
            var v34 int64 = 378
            var v35 uint64 = 42
            var v36 uint64 = (^(v35))
            _ = v34
            _ = v35
            _ = v36
        }
    case -9223372036854775808:
        for fi37 := 0; fi37 < 18; fi37++ {
            v38 := []int{-2147483648, -2147483648, v29, (fi37 >> ((v3) & 63))}
            if len(v38) > 0 {
                v38[0] = (v1 - v29)
            }
            _ = fi37
            _ = v38
        }
    case 225:
        var v39 bool = false
        _ = v39
    default:
        var v40 uint8 = 116
        _ = v40
    }
    v1 = v29
    var v41 int = (v1 >> (((v3 * v3)) & 63))
    v42 := GenS2{F0: ("" + "foo bar")}
    for fi43 := 0; fi43 < 19; fi43++ {
        var v44 bool = (158 > 1)
        v3 = v3
        v45 := []int{v29, 4294967295, (v1 | int(v3))}
        _ = fi43
        _ = v44
        _ = v45
    }
    var v46 int8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v46
}
