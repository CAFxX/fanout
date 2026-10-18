package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16; F1 uint8; F2 int64 }
type GenS3 struct { F0 int32; F1 bool; F2 int16 }
type GenS4 struct { F0 uint8; F1 int; F2 string }
type GenS5 struct { F0 int; F1 int8 }
func genFn0(p0 uint32) uint64 {
    v4 := []int{4294967295, -2147483648}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch 32767 {
    case -2147483648:
        for fi5 := 0; fi5 < 18; fi5++ {
            var v6 int8 = 100
            if false {
                var v7 bool = (255 < 1)
                _ = v7
            }
            _ = fi5
            _ = v6
        }
    case 0:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    _ = p0
    _ = v4
    return 684
}
func genFn1(p0 int, p1 string) uint16 {
    p1 = p1
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 int32 = -413
        v9 += v9
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    return 42
}

func main() {
    v1 := []uint32{2147483647, 786, 0}
    var v2 float64 = 1e10
    v3 := []int{2147483647}
    var v10 int32 = 42
    v11 := GenS1{F0: 2}
    var v12 string = "\n\t"
    fmt.Println(v2, v12, v10)
    v13 := GenS2{F0: 1000, F1: 188, F2: -127}
    v14 := GenS1{F0: 255}
    switch -347 {
    case 4294967295:
        if ((!(18446744073709551615 > 742)) || true) {
            for fi15 := 0; fi15 < 10; fi15++ {
                v2 = ((3.14159 + 0.1) - v2)
                v10 -= v10
                _ = fi15
            }
            v2 -= (v2 - float64((float32(v2) * (1.5 / 3.14159))))
        }
    default:
        v12 = (v12 + ((("\n\t" + v12) + v12) + "\n\t"))
    }
    if true {
        for fi16 := 0; fi16 < 20; fi16++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v17 := GenS3{F0: (v10 - (v10 ^ v10)), F1: (!false), F2: -264}
            _ = fi16
            _ = v17
        }
        for fi18 := 0; fi18 < 7; fi18++ {
            var v19 float32 = ((1.5 + (-2.25 / (-2.25 / -2.25))) * 3.14159)
            v12 = ("hello" + (((v12 + v12) + v12) + "foo bar"))
            _ = fi18
            _ = v19
        }
    }
    switch "hello" {
    case "hello":
        fmt.Println(v12, v2)
    case "a":
        switch (((v12 + (((v12 + v12) + ("hello" + "foo bar")) + v12)) + ((v12 + (v12 + "")) + ((v12 + ("hello" + "\n\t")) + ("foo bar" + "hello")))) + v12) {
        case "":
            var v20 uint16 = genFn1(-1, v12)
            _ = v20
        case "foo bar":
            for fi21 := 0; fi21 < 10; fi21++ {
                v12 = "\n\t"
                v22 := GenS2{F0: 566, F1: 180, F2: -1}
                if (false && false) {
                    v10 += v10
                }
                _ = fi21
                _ = v22
            }
        default:
            for fi23 := 0; fi23 < 13; fi23++ {
                for fi24 := 0; fi24 < 1; fi24++ {
                    v25 := GenS4{F0: 206, F1: -997, F2: v12}
                    var v26 uint = 291
                    fmt.Println(v10)
                    _ = fi24
                    _ = v25
                    _ = v26
                }
                for fi27 := 0; fi27 < 14; fi27++ {
                    var v28 bool = ((3.14159 / float32(v2)) <= (1e10 - -2.25))
                    var v29 uint8 = 103
                    _ = fi27
                    _ = v28
                    _ = v29
                }
                _ = fi23
            }
        }
    }
    var v30 string = ((v12 + (v12 + v12)) + ((v12 + ((v12 + (v12 + (v12 + v12))) + v12)) + (v12 + ("foo bar" + "a"))))
    if true {
        for fi31 := 0; fi31 < 4; fi31++ {
            switch fi31 {
            case -162:
                v32 := GenS4{F0: 2, F1: -2147483648, F2: "a"}
                _ = v32
            default:
                fmt.Println(v10)
            }
            switch fi31 {
            case -365:
                switch (fi31 + fi31) {
                case -566:
                    var v33 int64 = -9223372036854775808
                    _ = v33
                }
            case 312:
                for fi34 := 0; fi34 < 4; fi34++ {
                    v35 := GenS1{F0: 128}
                    v36 := GenS5{F0: -1, F1: -1}
                    v37 := GenS2{F0: -708, F1: 42, F2: 429}
                    _ = fi34
                    _ = v35
                    _ = v36
                    _ = v37
                }
            }
            for fi38 := 0; fi38 < 10; fi38++ {
                var v39 uint64 = 1
                var v40 uint32 = 0
                _ = fi38
                _ = v39
                _ = v40
            }
            _ = fi31
        }
    } else {
        var v41 int64 = 91
        _ = v41
    }
    var v42 uint8 = 4
    v3 = append(v3, 32767)
    if len(v1) > 0 {
        v1[0] = 2
    }
    var v43 uint16 = 65535
    var v44 uint16 = (v43 / ((v43) | 1))
    var v45 float64 = -2.25
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v30
    _ = v42
    _ = v43
    _ = v44
    _ = v45
}
