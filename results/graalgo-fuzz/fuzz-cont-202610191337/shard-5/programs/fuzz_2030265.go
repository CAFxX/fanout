package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 int8; F2 uint8 }
type GenS3 struct { F0 string; F1 int32; F2 int32 }
type GenS4 struct { F0 uint; F1 int16 }
type GenS5 struct { F0 int32; F1 string }
type GenS6 struct { F0 int16; F1 string }
func genFn0(p0 uint16, p1 int) uint {
    p1 -= 318
    fmt.Println(p0)
    var v6 string = (("" + "hello") + "")
    switch ("foo bar" + (v6 + v6)) {
    case "hello":
        var v7 int32 = 42
        _ = v7
    case "":
        v8 := []int{-429}
        _ = v8
    default:
        var v9 int32 = 1
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v6
    return 2
}
func genFn1() int8 {
    for fi10 := 0; fi10 < 11; fi10++ {
        fmt.Println(fi10)
        var v11 uint64 = 42
        v11 += 0
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 18; fi12++ {
        if ((!true) && (-510 == fi12)) {
            v13 := GenS1{F0: 15}
            v14 := []int{(^(fi12)), (fi12 ^ (fi12 - fi12)), 9223372036854775807, (fi12 ^ (fi12 & 2147483647))}
            for fi15 := 0; fi15 < 12; fi15++ {
                switch "" {
                case "a":
                    fmt.Println(fi12, fi15)
                }
                fmt.Println(fi12, fi12)
                _ = fi15
            }
            _ = v13
            _ = v14
        } else {
            for fi16 := 0; fi16 < 12; fi16++ {
                var v17 uint8 = 78
                fmt.Println(v17, fi16, fi16)
                v18 := []int{}
                _ = fi16
                _ = v17
                _ = v18
            }
            var v19 int16 = -693
            _ = v19
        }
        v20 := []int{fi12, -180}
        var v21 int8 = 0
        _ = fi12
        _ = v20
        _ = v21
    }
    for fi22 := 0; fi22 < 19; fi22++ {
        var v23 float64 = (3.14159 / 1.5)
        _ = fi22
        _ = v23
    }
    return -57
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint = 772
    v3 := make(map[string]bool)
    v4 := make(map[uint32]int)
    var v5 uint64 = 4294967295
    var v24 int16 = 179
    v25 := []int{}
    if true {
        for fi26 := 0; fi26 < 13; fi26++ {
            _ = len(v25)
            var v27 bool = (false || ("foo bar" <= "hello"))
            if len(v25) >= 2 {
                v25 = v25[1:len(v25)]
            }
            _ = fi26
            _ = v27
        }
        v24 += v24
        for fi28 := 0; fi28 < 15; fi28++ {
            v5 = v5
            fmt.Println(v5)
            switch ("a" + "") {
            case "hello":
                v29 := GenS2{F0: v24, F1: 42, F2: 127}
                _ = v29
            case "a":
                var v30 uint64 = v5
                _ = v30
            case "foo bar":
                v31 := GenS1{F0: 42}
                _ = v31
            }
            _ = fi28
        }
    } else {
        fmt.Println(v5, v5, v2)
        for fi32 := 0; fi32 < 1; fi32++ {
            v33 := GenS3{F0: "a", F1: 0, F2: -2147483648}
            _ = fi32
            _ = v33
        }
    }
    if (495 > int(v2)) {
        v24 = -1
        for fi34 := 0; fi34 < 6; fi34++ {
            if len(v25) >= 2 {
                v25 = v25[1:len(v25)]
            }
            _ = fi34
        }
    }
    v24 = v24
    for fi35 := 0; fi35 < 9; fi35++ {
        v25 = append(v25, int(v24))
        v5 = 0
        _ = fi35
    }
    if false {
        var v36 uint16 = 32767
        _ = v36
    }
    switch 128 {
    case 2147483647:
        fmt.Println(v2, v2, v24)
    default:
        var v37 bool = false
        _ = v37
    }
    var v38 bool = true
    v1[4294967295] = 2147483648
    v39 := v1[1]
    v40 := GenS4{F0: v2, F1: -1}
    v41 := GenS5{F0: 42, F1: (("\n\t" + "hello") + ("\n\t" + (("" + "foo bar") + ("hello" + "\n\t"))))}
    for fi42 := 0; fi42 < 2; fi42++ {
        v39 = uint32(v2)
        _ = fi42
    }
    var v43 float32 = 0.1
    var v44 bool = v38
    switch -579 {
    case 42:
        v45 := GenS6{F0: -823, F1: "hello"}
        _ = v45
    case -9223372036854775808:
        var v46 uint16 = 615
        _ = v46
    case -2147483648:
        fmt.Println(v44, v24)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v43
    _ = v44
    _ = v5
}
