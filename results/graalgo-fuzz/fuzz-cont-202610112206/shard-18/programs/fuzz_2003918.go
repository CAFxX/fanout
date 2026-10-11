package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int16; F1 int16 }
type GenS4 struct { F0 string; F1 string }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int64 }
func genFn0(p0 int16, p1 uint16) int {
    if false {
        if ("\n\t" > "foo bar") {
            if false {
                v6 := []int{}
                p1 += (p1 & p1)
                _ = v6
            }
            var v7 float64 = ((float64(1.5) + float64(1e10)) * 0.1)
            _ = v7
        }
        if (0.1 == -2.25) {
            var v8 uint16 = p1
            v9 := []int{9223372036854775807, -133, -663, -9223372036854775808}
            _ = v8
            _ = v9
        } else {
            if true {
                p1 += (^(p1))
                var v10 uint32 = 2147483648
                _ = v10
            }
        }
        if true {
            fmt.Println(p0)
            for fi11 := 0; fi11 < 15; fi11++ {
                v12 := GenS1{F0: "hello", F1: 823}
                var v13 uint64 = 4294967295
                _ = fi11
                _ = v12
                _ = v13
            }
        }
    }
    v14 := GenS2{F0: (!false)}
    _ = p0
    _ = p1
    _ = v14
    return int(p1)
}

func main() {
    v1 := []uint32{317}
    var v2 uint = 255
    var v3 string = ""
    v4 := make(map[int]int)
    v5 := make(map[string]uint32)
    switch v3 {
    case "\n\t":
        v3 = (v3 + "a")
    case "a":
        if false {
            fmt.Println(v3, v2)
            v2 += v2
        } else {
            v15 := GenS2{F0: (!(true && ((1 < 655) || (((v3 + ("" + "foo bar")) + "a") > (v3 + v3)))))}
            v3 = ""
            _ = v15
        }
    case "foo bar":
        v16 := GenS1{F0: ("hello" + ("\n\t" + (v3 + v3))), F1: 9223372036854775807}
        _ = v16
    }
    fmt.Println(v2, v3)
    fmt.Println(v2)
    for fi17 := 0; fi17 < 3; fi17++ {
        v18 := GenS3{F0: 0, F1: 42}
        var v19 bool = false
        switch fi17 {
        case 0:
            v2 = (v2 ^ (v2 + v2))
        }
        _ = fi17
        _ = v18
        _ = v19
    }
    v3 = ("hello" + v3)
    fmt.Println(v2)
    for fi20 := 0; fi20 < 18; fi20++ {
        v21 := GenS4{F0: ((v3 + (v3 + v3)) + (v3 + "\n\t")), F1: (v3 + v3)}
        v22 := GenS5{F0: fi20}
        _ = fi20
        _ = v21
        _ = v22
    }
    for fi23 := 0; fi23 < 15; fi23++ {
        v2 = 42
        for fi24 := 0; fi24 < 11; fi24++ {
            v3 = v3
            v3 = v3
            var v25 uint16 = uint16(v2)
            fmt.Println(v2, fi23, v3)
            _ = fi24
            _ = v25
        }
        _ = fi23
    }
    var v26 uint8 = 42
    v27 := v5[v3]
    switch genFn0(32767, 1) {
    case 4294967295:
        fmt.Println(v2)
    default:
        v27 = (^(v27))
    }
    v3 = ((v3 + ("" + (v3 + v3))) + "foo bar")
    for fi28 := 0; fi28 < 3; fi28++ {
        for fi29 := 0; fi29 < 12; fi29++ {
            var v30 bool = true
            if (42 != -91) {
                var v31 float32 = -2.25
                v32 := GenS1{F0: (v3 + ((v3 + "hello") + (((v3 + "a") + "") + "\n\t"))), F1: 1}
                v33 := GenS6{F0: -1}
                _ = v31
                _ = v32
                _ = v33
            } else {
                if (true || false) {
                    var v34 uint16 = 65535
                    _ = v34
                }
            }
            _ = fi29
            _ = v30
        }
        _ = fi28
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
