package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 uint, p1 uint8) uint32 {
    p0 += (p0 + (p0 ^ 4294967295))
    if true {
        for fi6 := 0; fi6 < 3; fi6++ {
            if false {
                var v7 int64 = 783
                _ = v7
            }
            p0 += (p0 + (^(p0)))
            var v8 int8 = 127
            _ = fi6
            _ = v8
        }
        p0 -= (p0 & p0)
    } else {
        var v9 float32 = -2.25
        switch -9223372036854775808 {
        case -111:
            var v10 int32 = 735
            _ = v10
        }
        for fi11 := 0; fi11 < 15; fi11++ {
            fmt.Println(v9)
            _ = fi11
        }
        _ = v9
    }
    v12 := []int{-1}
    switch ("foo bar" + "hello") {
    case "foo bar":
        var v13 float32 = 1.5
        _ = v13
    case "a":
        var v14 int16 = 1
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v12
    return 834
}
func genFn1(p0 string) int64 {
    v15 := GenS1{F0: 42, F1: -94}
    v16 := GenS2{F0: 255}
    v17 := []int{-1}
    _ = p0
    _ = v15
    _ = v16
    _ = v17
    return 42
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{("foo bar" + "\n\t")}
    var v3 float32 = (3.14159 / (1.5 * 1e10))
    v4 := []string{}
    var v5 int32 = -1
    v3 = 1e10
    var v18 uint16 = 102
    for fi19 := 0; fi19 < 18; fi19++ {
        if len(v4) > 0 {
            v4[0] = "hello"
        }
        var v20 uint16 = v18
        _ = fi19
        _ = v20
    }
    for fi21 := 0; fi21 < 14; fi21++ {
        v18 = (^(v18))
        var v22 float32 = (v3 + (-2.25 - v3))
        _ = len(v1)
        _ = fi21
        _ = v22
    }
    switch "hello" {
    case "a":
        v3 = (1.5 / (v3 * 1e10))
    case "\n\t":
        v23 := v1[65535]
        _ = v23
    }
    var v24 uint64 = 795
    if (0.1 <= -2.25) {
        v25 := GenS1{F0: v24, F1: 1}
        _ = v25
    } else {
        v26 := GenS2{F0: 255}
        for fi27 := 0; fi27 < 20; fi27++ {
            for fi28 := 0; fi28 < 13; fi28++ {
                var v29 uint = 0
                v30 := GenS2{F0: 2}
                _ = fi28
                _ = v29
                _ = v30
            }
            _ = fi27
        }
        for fi31 := 0; fi31 < 5; fi31++ {
            var v32 bool = true
            _ = fi31
            _ = v32
        }
        _ = v26
    }
    v33 := GenS2{F0: 234}
    switch 512 {
    case -667:
        if (1 == 105) {
            _ = len(v4)
            v18 = (v18 % ((v18) | 1))
        } else {
            v3 += (float32(-2.25) + v3)
        }
    case -834:
        if (0 <= -76) {
            fmt.Println(v18, v18, v24)
            v18 = 2
            if (-9223372036854775808 < 477) {
                v34 := GenS3{F0: 2}
                var v35 float32 = (v3 - 0.0)
                _ = v34
                _ = v35
            }
        } else {
            var v36 uint64 = v24
            v5 += -777
            for fi37 := 0; fi37 < 18; fi37++ {
                v3 = -2.25
                v38 := GenS3{F0: v24}
                var v39 bool = false
                _ = fi37
                _ = v38
                _ = v39
            }
            _ = v36
        }
    default:
        var v40 uint8 = 255
        _ = v40
    }
    for fi41 := 0; fi41 < 16; fi41++ {
        for fi42 := 0; fi42 < 3; fi42++ {
            v24 += (v24 & (v24 - 1))
            fmt.Println(v18, v24)
            var v43 uint8 = 128
            _ = fi42
            _ = v43
        }
        _ = fi41
    }
    if (!((true || (v3 >= v3)) || (true && true))) {
        switch ("\n\t" + "foo bar") {
        case "a":
            v24 = v24
        case "":
            v24 = v24
        case "foo bar":
            v44 := GenS2{F0: 48}
            _ = v44
        default:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
    } else {
        switch ((("a" + (("" + "a") + "hello")) + ("hello" + "a")) + "a") {
        case "\n\t":
            for fi45 := 0; fi45 < 19; fi45++ {
                for fi46 := 0; fi46 < 17; fi46++ {
                    _ = len(v4)
                    var v47 uint8 = 128
                    _ = fi46
                    _ = v47
                }
                _ = fi45
            }
        default:
            var v48 uint32 = genFn0(255, 0)
            _ = v48
        }
    }
    var v49 uint64 = v24
    var v50 uint16 = (v18 & 2)
    delete(v1, 0)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v18
    _ = v2
    _ = v24
    _ = v3
    _ = v33
    _ = v4
    _ = v49
    _ = v5
    _ = v50
}
