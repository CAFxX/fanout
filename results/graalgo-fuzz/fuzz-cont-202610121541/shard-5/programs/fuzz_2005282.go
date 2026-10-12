package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int; F1 int64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 bool; F1 string; F2 string }
type GenS5 struct { F0 uint32; F1 uint16 }
func genFn0() uint {
    v4 := GenS1{F0: "hello"}
    v5 := GenS2{F0: 42, F1: 0}
    switch "\n\t" {
    case "":
        v6 := []int{}
        _ = v6
    case "foo bar":
        for fi7 := 0; fi7 < 2; fi7++ {
            v8 := GenS3{F0: 848}
            if (false || (true && (!false))) {
                switch (fi7 | -697) {
                case 4294967295:
                    fmt.Println(fi7)
                case -557:
                    var v9 string = (("\n\t" + ("a" + ("" + ""))) + (("a" + "\n\t") + ("hello" + ("foo bar" + ""))))
                    _ = v9
                }
            }
            _ = fi7
            _ = v8
        }
    }
    _ = v4
    _ = v5
    return 255
}

func main() {
    var v1 string = ("foo bar" + "foo bar")
    var v2 uint32 = 1
    var v3 int64 = -1
    for fi10 := 0; fi10 < 5; fi10++ {
        v11 := []int{fi10, -137, fi10}
        _ = fi10
        _ = v11
    }
    v12 := []int{}
    if true {
        var v13 uint8 = 128
        _ = v13
    } else {
        var v14 float32 = (3.14159 + -2.25)
        var v15 int32 = 1
        for fi16 := 0; fi16 < 20; fi16++ {
            var v17 int64 = v3
            _ = len(v12)
            var v18 int16 = -1
            _ = fi16
            _ = v17
            _ = v18
        }
        _ = v14
        _ = v15
    }
    if true {
        if false {
            _ = len(v12)
        } else {
            if false {
                switch 30 {
                case 42:
                    fmt.Println(v3, v3, v3)
                default:
                    if (false || true) {
                        if (!(-715 < -1)) {
                            fmt.Println(v3, v1)
                        }
                        fmt.Println(v3, v1)
                    }
                }
                for fi19 := 0; fi19 < 15; fi19++ {
                    v20 := GenS2{F0: fi19, F1: v3}
                    _ = fi19
                    _ = v20
                }
            }
            v2 -= (v2 & (v2 * v2))
        }
        var v21 string = v1
        _ = v21
    } else {
        var v22 int = -1
        if (!true) {
            fmt.Println(v2, v3, v1)
            for fi23 := 0; fi23 < 16; fi23++ {
                v22 = (^(fi23))
                v12 = append(v12, v22)
                _ = fi23
            }
        } else {
            if (!true) {
                for fi24 := 0; fi24 < 2; fi24++ {
                    for fi25 := 0; fi25 < 18; fi25++ {
                        var v26 uint = genFn0()
                        if false {
                            v3 += v3
                        }
                        var v27 uint8 = 128
                        v28 := GenS1{F0: "foo bar"}
                        _ = fi25
                        _ = v26
                        _ = v27
                        _ = v28
                    }
                    _ = fi24
                }
                v12 = append(v12, (v22 - (v22 | (v22 * -2147483648))))
            }
            if len(v12) >= 2 {
                v12 = v12[1:len(v12)]
            }
        }
        _ = v22
    }
    v29 := GenS4{F0: (65535 < 654), F1: "\n\t", F2: v1}
    for fi30 := 0; fi30 < 4; fi30++ {
        if false {
            var v31 bool = (true || true)
            _ = v31
        }
        v2 = (v2 * v2)
        _ = fi30
    }
    var v32 int = -924
    for fi33 := 0; fi33 < 13; fi33++ {
        var v34 int32 = 1
        _ = fi33
        _ = v34
    }
    v35 := GenS5{F0: 2, F1: 2}
    var v36 bool = true
    for fi37 := 0; fi37 < 7; fi37++ {
        var v38 uint16 = uint16(v2)
        if (false && v36) {
            for fi39 := 0; fi39 < 10; fi39++ {
                v2 -= (v2 * (^(v2)))
                v2 = v2
                var v40 int16 = 929
                var v41 bool = (v38 != (v38 | 437))
                _ = fi39
                _ = v40
                _ = v41
            }
        }
        _ = fi37
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v12
    _ = v2
    _ = v29
    _ = v3
    _ = v32
    _ = v35
    _ = v36
}
