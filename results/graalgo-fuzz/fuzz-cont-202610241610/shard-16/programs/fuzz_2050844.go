package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 int8; F2 uint16 }
type GenS3 struct { F0 uint; F1 uint16 }
func genFn0(p0 uint32) int16 {
    if false {
        p0 = p0
        p0 += p0
        switch -2147483648 {
        case -809:
            if (100 <= 29) {
                p0 += p0
                p0 = (^(p0))
            } else {
                v5 := GenS1{F0: 2}
                for fi6 := 0; fi6 < 15; fi6++ {
                    v7 := []int{}
                    var v8 uint = 594
                    var v9 string = ("\n\t" + "")
                    _ = fi6
                    _ = v7
                    _ = v8
                    _ = v9
                }
                if (32768 < -9223372036854775808) {
                    fmt.Println(p0)
                    if ("\n\t" < "a") {
                        var v10 uint16 = 333
                        var v11 uint32 = (p0 / ((p0) | 1))
                        _ = v10
                        _ = v11
                    }
                } else {
                    for fi12 := 0; fi12 < 13; fi12++ {
                        v13 := []int{}
                        _ = fi12
                        _ = v13
                    }
                }
                _ = v5
            }
        case 9223372036854775807:
            v14 := GenS1{F0: 127}
            _ = v14
        case 128:
            var v15 bool = false
            _ = v15
        }
    }
    p0 = p0
    v16 := GenS1{F0: 0}
    p0 += p0
    _ = p0
    _ = v16
    return -939
}

func main() {
    v1 := []string{"", "\n\t", "\n\t", ""}
    v2 := []string{}
    v3 := make(map[uint32]uint32)
    v4 := make(map[string]uint32)
    v17 := GenS1{F0: 128}
    if (0 > -1) {
        var v18 uint16 = 32767
        for fi19 := 0; fi19 < 19; fi19++ {
            for fi20 := 0; fi20 < 10; fi20++ {
                if ((false && true) || true) {
                    if true {
                        var v21 uint64 = 0
                        _ = v21
                    }
                } else {
                    var v22 uint64 = 1
                    v2 = append(v2, "foo bar")
                    fmt.Println(v18, v18)
                    var v23 bool = true
                    _ = v22
                    _ = v23
                }
                fmt.Println(fi20)
                var v24 uint16 = (v18 * (v18 / (((v18 ^ v18)) | 1)))
                _ = fi20
                _ = v24
            }
            v18 += 378
            _ = fi19
        }
        v25 := GenS2{F0: 255, F1: 90, F2: v18}
        _ = v18
        _ = v25
    } else {
        var v26 int8 = 26
        _ = v26
    }
    for fi27 := 0; fi27 < 5; fi27++ {
        _ = len(v1)
        if ((0 >= -509) && true) {
            for fi28 := 0; fi28 < 16; fi28++ {
                var v29 string = "a"
                _ = fi28
                _ = v29
            }
            var v30 int = (-(fi27))
            _ = v30
        } else {
            var v31 string = (("foo bar" + "hello") + "hello")
            switch "\n\t" {
            case "a":
                if ((true && (v31 >= (("\n\t" + v31) + ""))) || (true && false)) {
                    v31 = v31
                }
            }
            v31 = v31
            _ = v31
        }
        for fi32 := 0; fi32 < 15; fi32++ {
            var v33 int32 = 499
            var v34 float32 = 0.1
            _ = fi32
            _ = v33
            _ = v34
        }
        _ = fi27
    }
    var v35 int8 = -82
    if false {
        v36 := v4["a"]
        var v37 float32 = 1.5
        _ = v36
        _ = v37
    } else {
        var v38 uint32 = 65535
        _ = v38
    }
    v35 = v35
    v35 = (^(v35))
    var v39 string = (((("\n\t" + ("hello" + (((("foo bar" + ("" + ("" + "\n\t"))) + "a") + "\n\t") + "foo bar"))) + ("\n\t" + "")) + "a") + "a")
    v40 := GenS3{F0: 1, F1: 1}
    v35 += (v35 ^ (^(v35)))
    var v41 int16 = 773
    for fi42 := 0; fi42 < 20; fi42++ {
        v43 := GenS2{F0: 1, F1: (v35 + (v35 << ((42) & 63))), F2: 727}
        _ = fi42
        _ = v43
    }
    if true {
        switch 127 {
        case 867:
            v41 += (^(v41))
        case 42:
            v41 += (^(v41))
        }
    } else {
        var v44 uint32 = 2147483648
        fmt.Println(v39, v41)
        if true {
            for fi45 := 0; fi45 < 20; fi45++ {
                v41 -= v41
                var v46 float64 = 1e10
                v46 = -2.25
                _ = fi45
                _ = v46
            }
            var v47 uint32 = 294
            _ = v47
        } else {
            var v48 float64 = float64(0.0)
            var v49 int16 = v41
            _ = v48
            _ = v49
        }
        _ = v44
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v35
    _ = v39
    _ = v4
    _ = v40
    _ = v41
}
