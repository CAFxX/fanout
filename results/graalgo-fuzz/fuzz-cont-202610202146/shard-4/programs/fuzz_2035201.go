package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int; F1 uint16; F2 int16 }
type GenS4 struct { F0 bool; F1 bool }
type GenS5 struct { F0 uint16 }
func genFn0() int8 {
    var v7 float64 = 0.0
    switch "a" {
    case "foo bar":
        var v8 int8 = 69
        _ = v8
    case "":
        switch 1 {
        case 4294967295:
            for fi9 := 0; fi9 < 12; fi9++ {
                fmt.Println(fi9, fi9)
                if (100 < 0) {
                    var v10 uint = 255
                    _ = v10
                } else {
                    v11 := []int{}
                    _ = v11
                }
                _ = fi9
            }
        }
    case "a":
        v12 := []int{}
        _ = v12
    default:
        v7 -= (-2.25 * float64(float32(float64(1.5))))
    }
    _ = v7
    return -1
}
func genFn1() int32 {
    var v13 int16 = -32768
    switch "foo bar" {
    case "hello":
        for fi14 := 0; fi14 < 7; fi14++ {
            var v15 uint32 = 2
            _ = fi14
            _ = v15
        }
    }
    v13 -= v13
    _ = v13
    return -2147483648
}

func main() {
    var v1 int16 = -1
    var v2 int64 = -698
    var v3 string = "\n\t"
    v4 := []string{("\n\t" + "hello"), "a", "a", (v3 + v3)}
    var v5 int64 = (v2 / (((^(v2))) | 1))
    v6 := []int{779, -271}
    for fi16 := 0; fi16 < 19; fi16++ {
        v17 := GenS1{F0: 65535, F1: 42}
        if len(v4) > 0 {
            v4[0] = v3
        }
        _ = fi16
        _ = v17
    }
    var v18 uint8 = 165
    v2 -= v5
    for fi19 := 0; fi19 < 6; fi19++ {
        if len(v6) > 0 {
            v6[0] = fi19
        }
        _ = len(v4)
        _ = fi19
    }
    switch 127 {
    case -1:
        v20 := GenS2{F0: false}
        _ = v20
    case 9223372036854775807:
        v1 -= (v1 & 1)
    case 127:
        var v21 int32 = -1
        _ = v21
    default:
        if ((5 > 255) && false) {
            if len(v6) > 0 {
                v6[0] = 9223372036854775807
            }
        } else {
            var v22 string = (v3 + "\n\t")
            for fi23 := 0; fi23 < 20; fi23++ {
                _ = len(v4)
                var v24 uint8 = (^(v18))
                _ = fi23
                _ = v24
            }
            _ = v22
        }
    }
    fmt.Println(v18)
    var v25 int = -1
    switch (v3 + "foo bar") {
    case "a":
        var v26 uint32 = 1
        _ = v26
    case "\n\t":
        v27 := GenS2{F0: false}
        _ = v27
    }
    if true {
        if true {
            v28 := GenS3{F0: v25, F1: 32767, F2: v1}
            var v29 uint = 658
            _ = v28
            _ = v29
        }
    }
    if (v2 <= (v5 ^ 1)) {
        switch v25 {
        case 42:
            v30 := GenS4{F0: false, F1: false}
            _ = v30
        case -839:
            switch (v25 * v25) {
            case -9223372036854775808:
                if (!false) {
                    var v31 int64 = (v5 << ((1) & 63))
                    v2 = 959
                    _ = v31
                } else {
                    if len(v6) > 0 {
                        v6[0] = v25
                    }
                    var v32 uint = 2
                    _ = v32
                }
            case -487:
                var v33 uint32 = 65535
                _ = v33
            default:
                _ = len(v6)
            }
        }
        v5 -= 1
        for fi34 := 0; fi34 < 18; fi34++ {
            var v35 float32 = 1e10
            v2 = v2
            fmt.Println(v5)
            _ = fi34
            _ = v35
        }
    } else {
        v36 := GenS5{F0: 568}
        _ = len(v6)
        _ = v36
    }
    switch v25 {
    case 2147483647:
        v37 := GenS5{F0: 952}
        _ = v37
    }
    for fi38 := 0; fi38 < 20; fi38++ {
        var v39 float32 = ((1e10 - 0.1) / 3.14159)
        switch (fi38 - fi38) {
        case 0:
            fmt.Println(fi38, v25)
        case 251:
            var v40 int32 = 1
            _ = v40
        case 2147483647:
            var v41 string = v3
            _ = v41
        }
        _ = fi38
        _ = v39
    }
    v5 = int64(v25)
    _ = len(v4)
    for fi42 := 0; fi42 < 14; fi42++ {
        var v43 int8 = 37
        switch v3 {
        case "":
            v5 += (v2 >> ((0) & 63))
        case "\n\t":
            var v44 int = 9223372036854775807
            _ = v44
        }
        _ = fi42
        _ = v43
    }
    v25 -= v25
    v5 = (v5 + v5)
    for fi45 := 0; fi45 < 20; fi45++ {
        for fi46 := 0; fi46 < 10; fi46++ {
            var v47 uint8 = 1
            var v48 bool = false
            _ = fi46
            _ = v47
            _ = v48
        }
        _ = fi45
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
