package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool; F1 uint; F2 uint }
type GenS3 struct { F0 uint64; F1 uint; F2 bool }
type GenS4 struct { F0 uint8; F1 string; F2 uint64 }
func genFn0(p0 int16) int16 {
    var v5 bool = (false && ((p0 > 593) && false))
    var v6 uint = 1
    _ = p0
    _ = v5
    _ = v6
    return -618
}
func genFn1() int32 {
    var v7 float64 = 0.0
    if true {
        for fi8 := 0; fi8 < 20; fi8++ {
            v9 := GenS1{F0: -128}
            _ = fi8
            _ = v9
        }
        v10 := GenS1{F0: -109}
        var v11 string = "hello"
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 14; fi12++ {
        var v13 float32 = ((0.0 / 1.0) + (-2.25 / 1.0))
        v14 := []int{}
        if (2 < 24) {
            switch "a" {
            case "a":
                if false {
                    var v15 uint64 = 2
                    _ = v15
                } else {
                    if false {
                        if (255 >= 65535) {
                            v16 := GenS2{F0: false, F1: 42, F2: 990}
                            _ = v16
                        } else {
                            v17 := GenS3{F0: 1, F1: 1, F2: true}
                            v18 := GenS3{F0: 2, F1: 65535, F2: false}
                            _ = v17
                            _ = v18
                        }
                    } else {
                        var v19 float64 = float64(1e10)
                        var v20 uint64 = 0
                        _ = v19
                        _ = v20
                    }
                }
            }
        }
        _ = fi12
        _ = v13
        _ = v14
    }
    var v21 uint16 = 0
    _ = v21
    _ = v7
    return 2147483647
}

func main() {
    v1 := make(map[string]bool)
    v2 := []uint32{2, 1, 2147483647}
    var v3 int64 = 755
    var v4 string = ("\n\t" + "foo bar")
    v22 := GenS2{F0: true, F1: 18446744073709551615, F2: 318}
    if false {
        var v23 uint16 = 168
        var v24 int8 = 1
        _ = v23
        _ = v24
    } else {
        for fi25 := 0; fi25 < 15; fi25++ {
            v4 = "a"
            v1[v4] = false
            _ = fi25
        }
        fmt.Println(v3)
        var v26 int32 = 1
        _ = v26
    }
    if len(v2) > 0 {
        v2[0] = 266
    }
    if ("\n\t" == "\n\t") {
        var v27 float64 = float64(((0.1 - ((1.5 + float32(1.5)) / 0.1)) - 0.1))
        if (true && (-2147483648 < 42)) {
            for fi28 := 0; fi28 < 2; fi28++ {
                v4 = "foo bar"
                fmt.Println(v3, v3, v27)
                for fi29 := 0; fi29 < 19; fi29++ {
                    v30 := GenS2{F0: true, F1: 255, F2: 1}
                    _ = fi29
                    _ = v30
                }
                _ = fi28
            }
            v4 = ("\n\t" + ((v4 + (((v4 + ((((((((v4 + "hello") + "foo bar") + "hello") + v4) + ("\n\t" + ((("a" + (v4 + ("hello" + ("hello" + v4)))) + v4) + ((((v4 + ("\n\t" + v4)) + (v4 + v4)) + v4) + v4)))) + v4) + v4) + ("\n\t" + "hello"))) + v4) + (v4 + v4))) + v4))
        } else {
            for fi31 := 0; fi31 < 13; fi31++ {
                for fi32 := 0; fi32 < 10; fi32++ {
                    var v33 uint16 = 326
                    fmt.Println(v3, v4, v27)
                    _ = fi32
                    _ = v33
                }
                if (!(!false)) {
                    v34 := GenS2{F0: ((0.1 / float32((v27 - 1e10))) == 0.0), F1: 65535, F2: 4294967295}
                    v4 = ""
                    _ = v34
                }
                _ = fi31
            }
            v35 := v1[(("foo bar" + "foo bar") + "a")]
            _ = v35
        }
        var v36 string = ("hello" + ((v4 + "foo bar") + ""))
        _ = v27
        _ = v36
    }
    for fi37 := 0; fi37 < 17; fi37++ {
        v3 += v3
        var v38 int32 = 2147483647
        _ = fi37
        _ = v38
    }
    v4 = v4
    if false {
        for fi39 := 0; fi39 < 1; fi39++ {
            v40 := GenS4{F0: 2, F1: v4, F2: 0}
            _ = fi39
            _ = v40
        }
    } else {
        v41 := GenS1{F0: 100}
        v3 = (v3 % ((v3) | 1))
        _ = v41
    }
    v3 += (v3 + (v3 << ((42) & 63)))
    var v42 uint16 = 0
    switch 963 {
    case 9223372036854775807:
        if (false || false) {
            if len(v2) > 0 {
                v2[0] = 2147483648
            }
            v3 = 0
        } else {
            if ((v42 >> ((276) & 63)) < 309) {
                if true {
                    var v43 uint = 255
                    fmt.Println(v4)
                    _ = v43
                }
                v44 := GenS3{F0: 437, F1: 18446744073709551615, F2: (!(true && (32768 >= 32768)))}
                _ = v44
            }
        }
    case 128:
        var v45 uint = 598
        _ = v45
    }
    var v46 uint8 = 152
    v47 := GenS3{F0: 1, F1: 255, F2: false}
    if false {
        var v48 int = 9223372036854775807
        _ = v48
    }
    fmt.Println(v3)
    var v49 uint = 18446744073709551615
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v46:", v46)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v42
    _ = v46
    _ = v47
    _ = v49
}
