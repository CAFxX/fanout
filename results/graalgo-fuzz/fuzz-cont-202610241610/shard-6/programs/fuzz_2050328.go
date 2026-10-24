package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 int; F1 bool }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int8; F1 int64; F2 uint16 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint64 }

func main() {
    var v1 bool = (1 == 65535)
    var v2 float32 = -2.25
    var v3 uint8 = 255
    var v4 string = "\n\t"
    v3 -= (v3 / ((127) | 1))
    v5 := []int{-990}
    if v1 {
        v6 := GenS1{F0: 42, F1: 228}
        if len(v5) > 0 {
            v5[0] = -9223372036854775808
        }
        for fi7 := 0; fi7 < 6; fi7++ {
            fmt.Println(v1, v1, v2)
            _ = fi7
        }
        _ = v6
    }
    var v8 float32 = v2
    if (!(32767 <= -182)) {
        var v9 bool = false
        for fi10 := 0; fi10 < 20; fi10++ {
            for fi11 := 0; fi11 < 18; fi11++ {
                v3 = v3
                switch (fi11 << ((42) & 63)) {
                case -802:
                    if ((v9 || (42 == 4294967295)) && (65535 != 2)) {
                        v8 += (1e10 - 0.0)
                    }
                }
                v12 := GenS1{F0: 2147483647, F1: 177}
                var v13 float32 = v8
                if (255 == 887) {
                    v14 := GenS1{F0: 4294967295, F1: 827}
                    _ = v14
                }
                _ = fi11
                _ = v12
                _ = v13
            }
            v3 = v3
            if false {
                var v15 uint = 255
                switch (fi10 + 9223372036854775807) {
                case 1:
                    v4 = v4
                }
                if len(v5) > 0 {
                    v5[0] = (fi10 | 128)
                }
                _ = v15
            } else {
                var v16 bool = (v9 && true)
                for fi17 := 0; fi17 < 10; fi17++ {
                    if len(v5) > 0 {
                        v5[0] = fi17
                    }
                    var v18 int8 = 102
                    if true {
                        fmt.Println(v9, fi17, v16)
                    }
                    _ = fi17
                    _ = v18
                }
                _ = len(v5)
                _ = v16
            }
            _ = fi10
        }
        v5 = append(v5, 127)
        _ = v9
    } else {
        v19 := GenS1{F0: 94, F1: 0}
        var v20 uint8 = (^(v3))
        v21 := GenS2{F0: 9223372036854775807, F1: v1}
        _ = v19
        _ = v20
        _ = v21
    }
    for fi22 := 0; fi22 < 1; fi22++ {
        if v1 {
            v5 = append(v5, fi22)
            if v1 {
                if len(v5) > 0 {
                    v5[0] = fi22
                }
            } else {
                if len(v5) > 0 {
                    v5[0] = int(v3)
                }
                if false {
                    v4 = "a"
                    for fi23 := 0; fi23 < 10; fi23++ {
                        fmt.Println(v3, fi22)
                        _ = fi23
                    }
                    v24 := GenS3{F0: -1}
                    _ = v24
                } else {
                    v25 := GenS1{F0: 2147483648, F1: 2}
                    v3 += v3
                    if false {
                        for fi26 := 0; fi26 < 6; fi26++ {
                            v8 += v2
                            v27 := GenS3{F0: 32767}
                            _ = fi26
                            _ = v27
                        }
                    } else {
                        var v28 float64 = (1.5 * float64(v2))
                        v29 := GenS4{F0: -1, F1: 1, F2: 42}
                        var v30 uint16 = 32767
                        _ = v28
                        _ = v29
                        _ = v30
                    }
                    _ = v25
                }
                for fi31 := 0; fi31 < 13; fi31++ {
                    switch "a" {
                    case "a":
                        var v32 int = (fi22 % ((-393) | 1))
                        _ = v32
                    }
                    var v33 float64 = (1e10 / 0.1)
                    _ = fi31
                    _ = v33
                }
            }
            v34 := GenS5{F0: ((^(v3)) != (v3 | 0))}
            _ = v34
        }
        _ = fi22
    }
    var v35 bool = true
    v1 = true
    _ = len(v5)
    v36 := GenS6{F0: 475}
    if (-2147483648 < 1) {
        for fi37 := 0; fi37 < 14; fi37++ {
            v38 := GenS2{F0: fi37, F1: v35}
            fmt.Println(v1)
            v3 -= v3
            _ = fi37
            _ = v38
        }
        switch (v4 + "hello") {
        case "":
            var v39 uint8 = v3
            _ = v39
        }
    }
    v4 = ("" + "hello")
    if ((100 == -1) && true) {
        v35 = v35
    }
    var v40 int = 128
    var v41 uint16 = 77
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v8
}
