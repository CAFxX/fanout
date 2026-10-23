package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64; F1 uint; F2 uint8 }
type GenS4 struct { F0 string; F1 int64 }
func genFn0(p0 int16, p1 int8) int8 {
    fmt.Println(p0)
    var v4 uint8 = 127
    for fi5 := 0; fi5 < 17; fi5++ {
        v6 := []int{(-(fi5)), (^(fi5))}
        var v7 uint = 1
        _ = fi5
        _ = v6
        _ = v7
    }
    if (true || false) {
        v4 = (v4 << ((212) & 63))
        if true {
            var v8 int = -878
            var v9 int8 = (-(p1))
            _ = v8
            _ = v9
        }
    }
    var v10 int = -8
    _ = p0
    _ = p1
    _ = v10
    _ = v4
    return (p1 + int8(v4))
}
func genFn1(p0 bool) uint64 {
    var v11 int8 = 42
    switch "\n\t" {
    case "\n\t":
        if (1 < 1) {
            var v12 float32 = 0.1
            var v13 uint = 1
            _ = v12
            _ = v13
        } else {
            v14 := GenS1{F0: 32768}
            _ = v14
        }
    case "":
        fmt.Println(v11, p0)
    }
    v15 := GenS2{F0: ((42 < 1) || p0)}
    for fi16 := 0; fi16 < 13; fi16++ {
        v17 := GenS2{F0: p0}
        p0 = p0
        if p0 {
            switch "a" {
            case "a":
                fmt.Println(fi16)
            case "hello":
                p0 = p0
            case "foo bar":
                v18 := []int{fi16, (^(fi16)), (fi16 >> ((2) & 63))}
                _ = v18
            }
            v19 := GenS1{F0: (fi16 ^ fi16)}
            _ = v19
        } else {
            v20 := GenS3{F0: 852, F1: 18446744073709551615, F2: 144}
            switch (fi16 >> ((1) & 63)) {
            case 1:
                for fi21 := 0; fi21 < 20; fi21++ {
                    var v22 uint64 = 1
                    _ = fi21
                    _ = v22
                }
            case -1:
                for fi23 := 0; fi23 < 12; fi23++ {
                    fmt.Println(fi16)
                    _ = fi23
                }
            case -860:
                var v24 int8 = v11
                _ = v24
            }
            p0 = true
            _ = v20
        }
        _ = fi16
        _ = v17
    }
    if p0 {
        v25 := GenS4{F0: "a", F1: 37}
        for fi26 := 0; fi26 < 18; fi26++ {
            v27 := GenS4{F0: (("a" + "\n\t") + (("\n\t" + (("hello" + "") + "")) + "\n\t")), F1: 1}
            _ = fi26
            _ = v27
        }
        _ = v25
    }
    _ = p0
    _ = v11
    _ = v15
    return 18446744073709551615
}

func main() {
    var v1 int32 = -2147483648
    v2 := make(map[uint32]string)
    var v3 int8 = 25
    v28 := []int{32768}
    for fi29 := 0; fi29 < 11; fi29++ {
        v1 = (v1 - (v1 << ((4294967295) & 63)))
        _ = fi29
    }
    v1 += 0
    fmt.Println(v3, v3)
    if (255 < 0) {
        v30 := GenS2{F0: (("foo bar" >= ("foo bar" + (("foo bar" + "") + "\n\t"))) && (v3 == v3))}
        v1 = (v1 - v1)
        v31 := v2[228]
        _ = v30
        _ = v31
    }
    delete(v2, 65535)
    if false {
        for fi32 := 0; fi32 < 14; fi32++ {
            v1 += (v1 / (((v1 << ((18446744073709551615) & 63))) | 1))
            for fi33 := 0; fi33 < 10; fi33++ {
                v3 -= (^(v3))
                v34 := GenS4{F0: (("\n\t" + "\n\t") + ("hello" + "a")), F1: 1}
                var v35 float64 = 0.0
                var v36 uint8 = 2
                _ = fi33
                _ = v34
                _ = v35
                _ = v36
            }
            var v37 uint = 843
            _ = fi32
            _ = v37
        }
        for fi38 := 0; fi38 < 15; fi38++ {
            fmt.Println(v3)
            v39 := GenS3{F0: 174, F1: 18446744073709551615, F2: 229}
            for fi40 := 0; fi40 < 9; fi40++ {
                var v41 uint = 42
                fmt.Println(fi38)
                var v42 uint8 = 42
                _ = fi40
                _ = v41
                _ = v42
            }
            _ = fi38
            _ = v39
        }
    } else {
        v3 = (v3 >> ((4294967295) & 63))
        var v43 int32 = (^(v1))
        _ = v43
    }
    var v44 float32 = ((float32(0.0) - 0.1) / 1.0)
    v1 -= v1
    if (1 >= 0) {
        var v45 uint32 = 0
        for fi46 := 0; fi46 < 7; fi46++ {
            switch "a" {
            case "hello":
                v45 -= v45
            case "a":
                var v47 uint32 = (v45 ^ 4294967295)
                _ = v47
            case "foo bar":
                var v48 float64 = float64((0.0 * v44))
                _ = v48
            }
            v1 = v1
            _ = fi46
        }
        _ = v45
    }
    for fi49 := 0; fi49 < 19; fi49++ {
        switch int(v3) {
        case 4294967295:
            v3 = v3
        case -9223372036854775808:
            v2[4294967295] = (("hello" + "a") + "\n\t")
        default:
            for fi50 := 0; fi50 < 8; fi50++ {
                v2[128] = "a"
                fmt.Println(fi49, v44, v3)
                _ = fi50
            }
        }
        _ = fi49
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v44
}
