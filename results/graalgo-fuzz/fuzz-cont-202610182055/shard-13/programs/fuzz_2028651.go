package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 int8; F1 uint32 }
type GenS4 struct { F0 string; F1 uint; F2 bool }
func genFn0() int {
    var v4 int64 = -111
    for fi5 := 0; fi5 < 8; fi5++ {
        if false {
            var v6 float64 = 1.5
            _ = v6
        }
        fmt.Println(fi5, v4)
        v7 := GenS1{F0: fi5}
        _ = fi5
        _ = v7
    }
    for fi8 := 0; fi8 < 18; fi8++ {
        v9 := []int{fi8, fi8}
        var v10 bool = (fi8 == -441)
        var v11 string = "\n\t"
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    var v12 float32 = ((float32((float64(((1.5 / 1.0) * 0.1)) * float64(1e10))) - 3.14159) - -2.25)
    v13 := GenS1{F0: -2147483648}
    _ = v12
    _ = v13
    _ = v4
    return 9223372036854775807
}

func main() {
    var v1 uint16 = 255
    v2 := make(map[uint32]bool)
    v3 := []int{}
    v1 += (^(v1))
    v14 := GenS1{F0: -1}
    for fi15 := 0; fi15 < 11; fi15++ {
        switch ("\n\t" + "\n\t") {
        case "":
            var v16 float32 = 0.0
            _ = v16
        case "foo bar":
            v17 := GenS1{F0: (fi15 % (((fi15 << ((913) & 63))) | 1))}
            _ = v17
        case "a":
            var v18 uint16 = v1
            _ = v18
        }
        if (((false && (v1 > 2)) && (0 > 65535)) && false) {
            v1 = v1
            v19 := GenS1{F0: 32767}
            fmt.Println(v1, fi15)
            _ = v19
        }
        _ = fi15
    }
    fmt.Println(v1)
    if len(v3) > 0 {
        v3[0] = 9223372036854775807
    }
    v20 := GenS1{F0: 252}
    var v21 uint64 = 18446744073709551615
    v22 := GenS1{F0: 9223372036854775807}
    v1 += (v1 >> ((550) & 63))
    v2[65535] = ((!(1 == 0)) || ((float64(0.1) >= float64(-2.25)) || true))
    if true {
        if false {
            var v23 uint32 = 365
            v3 = append(v3, 1)
            _ = v23
        } else {
            if (v21 >= (^(v21))) {
                var v24 uint16 = 2
                _ = len(v3)
                if len(v3) > 0 {
                    v3[0] = -934
                }
                _ = v24
            }
        }
    } else {
        var v25 float32 = -2.25
        _ = v25
    }
    v26 := GenS1{F0: 4294967295}
    fmt.Println(v1)
    var v27 uint32 = 65535
    if false {
        if (-32768 > -394) {
            fmt.Println(v1)
            v1 = (v1 - v1)
        }
        var v28 uint32 = v27
        _ = v28
    } else {
        var v29 float64 = 0.1
        v30 := GenS1{F0: 2147483647}
        _ = v29
        _ = v30
    }
    if (18446744073709551615 <= 578) {
        switch genFn0() {
        case 4294967295:
            if false {
                switch (("hello" + "\n\t") + ("foo bar" + "hello")) {
                case "hello":
                    v31 := GenS1{F0: genFn0()}
                    _ = v31
                case "":
                    v27 = (v27 * (v27 % (((v27 | (v27 | 1))) | 1)))
                case "\n\t":
                    var v32 float64 = 3.14159
                    _ = v32
                default:
                    var v33 int = 85
                    _ = v33
                }
            } else {
                for fi34 := 0; fi34 < 15; fi34++ {
                    var v35 float64 = 0.0
                    for fi36 := 0; fi36 < 12; fi36++ {
                        v27 += v27
                        v3 = append(v3, (fi36 << ((42) & 63)))
                        var v37 uint8 = 0
                        _ = fi36
                        _ = v37
                    }
                    if (42 < -128) {
                        v2[v27] = true
                        switch fi34 {
                        case -622:
                            switch fi34 {
                            case -2147483648:
                                v38 := GenS2{F0: 1, F1: (((("\n\t" + "foo bar") + "a") + "foo bar") + "a")}
                                _ = v38
                            }
                        case -1:
                            v39 := GenS3{F0: 127, F1: (v27 << ((1) & 63))}
                            _ = v39
                        }
                    }
                    _ = fi34
                    _ = v35
                }
                fmt.Println(v1, v21, v1)
            }
        default:
            for fi40 := 0; fi40 < 15; fi40++ {
                v41 := GenS3{F0: -91, F1: 65535}
                fmt.Println(v1, fi40)
                _ = fi40
                _ = v41
            }
        }
    } else {
        v42 := GenS1{F0: -901}
        _ = v42
    }
    for fi43 := 0; fi43 < 13; fi43++ {
        v1 = (v1 ^ 815)
        v44 := GenS4{F0: ((("" + "") + ("a" + ((("a" + "hello") + ("\n\t" + ("\n\t" + ""))) + "\n\t"))) + "foo bar"), F1: 0, F2: true}
        var v45 uint = 783
        _ = fi43
        _ = v44
        _ = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v26
    _ = v27
    _ = v3
}
