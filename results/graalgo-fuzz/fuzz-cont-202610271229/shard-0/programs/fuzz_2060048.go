package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int; F1 uint; F2 bool }
func genFn0(p0 string, p1 uint8, p2 uint64) bool {
    v5 := []int{-413}
    v6 := GenS1{F0: -136, F1: (^(p2))}
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return (p1 > 127)
}
func genFn1(p0 int32, p1 int8) int16 {
    switch "\n\t" {
    case "\n\t":
        var v7 float32 = -2.25
        _ = v7
    default:
        v8 := []int{4294967295, -104, 9223372036854775807}
        _ = v8
    }
    p0 += -2147483648
    _ = p0
    _ = p1
    return 733
}

func main() {
    var v1 uint = 2
    var v2 int32 = 1
    var v3 bool = ((v2 % (((v2 % (((v2 | -1)) | 1))) | 1)) == 2147483647)
    var v4 string = "\n\t"
    v9 := GenS1{F0: 128, F1: 129}
    fmt.Println(v2, v2, v2)
    var v10 uint32 = 436
    switch 2147483647 {
    case -275:
        v4 = v4
    case -2147483648:
        v2 = 176
    case -189:
        var v11 int64 = int64(v10)
        _ = v11
    }
    var v12 uint8 = 42
    if (false && v3) {
        fmt.Println(v3, v10, v12)
        switch 127 {
        case 2147483647:
            v13 := []int{127, 314, 128, -9223372036854775808}
            _ = v13
        }
        if (v2 >= v2) {
            v14 := GenS2{F0: 42}
            if true {
                v2 = v2
                v15 := GenS2{F0: int16(v12)}
                _ = v15
            } else {
                v16 := []int{0, 9223372036854775807, 127, -205}
                _ = v16
            }
            v17 := []int{int(v1), -454, 987, 4294967295}
            _ = v14
            _ = v17
        }
    } else {
        switch "a" {
        case "\n\t":
            var v18 int8 = 1
            _ = v18
        case "":
            var v19 int = -1
            _ = v19
        case "hello":
            v20 := []int{-1, 32768}
            _ = v20
        }
        for fi21 := 0; fi21 < 17; fi21++ {
            for fi22 := 0; fi22 < 20; fi22++ {
                fmt.Println(v4)
                var v23 uint = v1
                _ = fi22
                _ = v23
            }
            var v24 int = -1
            _ = fi21
            _ = v24
        }
    }
    var v25 uint32 = (v10 << (((v1 * (^(v1)))) & 63))
    for fi26 := 0; fi26 < 16; fi26++ {
        v27 := GenS3{F0: fi26, F1: v1, F2: true}
        var v28 float64 = (-2.25 - 1.5)
        if v3 {
            fmt.Println(v25, fi26, v28)
            fmt.Println(v25, v10)
            v1 += v1
        }
        _ = fi26
        _ = v27
        _ = v28
    }
    v2 += (v2 & v2)
    if true {
        v29 := []int{-805, 32767, 4294967295, 0}
        if len(v29) >= 2 {
            v29 = v29[1:len(v29)]
        }
        for fi30 := 0; fi30 < 5; fi30++ {
            switch ("\n\t" + v4) {
            case "a":
                var v31 float32 = 0.0
                _ = v31
            case "hello":
                fmt.Println(v2, fi30)
            }
            _ = fi30
        }
        _ = v29
    } else {
        var v32 uint8 = 0
        fmt.Println(v10)
        if false {
            if (0.1 < (0.1 / 3.14159)) {
                v12 = (v12 >> (((v1 * v1)) & 63))
                for fi33 := 0; fi33 < 19; fi33++ {
                    v34 := []int{(fi33 | fi33)}
                    var v35 int8 = 100
                    _ = fi33
                    _ = v34
                    _ = v35
                }
                v36 := GenS2{F0: 1}
                _ = v36
            }
            v25 = (v25 - 0)
            for fi37 := 0; fi37 < 13; fi37++ {
                var v38 uint = v1
                _ = fi37
                _ = v38
            }
        } else {
            var v39 int = -1
            var v40 int = (^(v39))
            _ = v39
            _ = v40
        }
        _ = v32
    }
    v10 -= (v10 & 2147483648)
    v41 := GenS2{F0: 1000}
    for fi42 := 0; fi42 < 20; fi42++ {
        v12 = v12
        _ = fi42
    }
    v43 := []int{-875}
    switch ("a" + v4) {
    case "\n\t":
        var v44 uint = 65535
        _ = v44
    case "foo bar":
        var v45 int32 = v2
        _ = v45
    case "hello":
        for fi46 := 0; fi46 < 19; fi46++ {
            for fi47 := 0; fi47 < 18; fi47++ {
                var v48 int64 = 0
                var v49 float32 = (1e10 - float32((float64((float32(3.14159) + 1.5)) + (float64(3.14159) - 1e10))))
                _ = fi47
                _ = v48
                _ = v49
            }
            _ = fi46
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v43:", len(v43), v43)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v41
    _ = v43
    _ = v9
}
