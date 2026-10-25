package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 string }
type GenS2 struct { F0 uint8; F1 bool; F2 bool }
type GenS3 struct { F0 int; F1 string; F2 string }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int; F1 int64; F2 int64 }
type GenS6 struct { F0 int16; F1 uint }
func genFn0(p0 uint64, p1 int16) uint16 {
    var v4 uint64 = 0
    v5 := GenS1{F0: "", F1: p1, F2: ""}
    if false {
        var v6 float64 = 3.14159
        _ = v6
    } else {
        v7 := GenS2{F0: 255, F1: true, F2: false}
        p1 = -648
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 65535
}
func genFn1() int16 {
    v8 := []int{127}
    if true {
        var v9 int8 = -1
        fmt.Println(v9)
        v8 = append(v8, -9223372036854775808)
        _ = v9
    }
    if (!false) {
        var v10 string = "hello"
        for fi11 := 0; fi11 < 1; fi11++ {
            fmt.Println(v10)
            _ = fi11
        }
        _ = v10
    }
    for fi12 := 0; fi12 < 10; fi12++ {
        fmt.Println(fi12)
        var v13 uint16 = 209
        _ = fi12
        _ = v13
    }
    var v14 uint64 = 0
    _ = v14
    _ = v8
    return -32768
}

func main() {
    var v1 float64 = (0.1 + ((float64(1e10) / 3.14159) / 1e10))
    var v2 int64 = 1
    v3 := make(map[uint32]string)
    v15 := GenS3{F0: 0, F1: "foo bar", F2: "hello"}
    if false {
        for fi16 := 0; fi16 < 13; fi16++ {
            v1 += v1
            for fi17 := 0; fi17 < 4; fi17++ {
                switch 32767 {
                case 32767:
                    v1 = (3.14159 * v1)
                case 350:
                    var v18 int = (fi16 << ((42) & 63))
                    _ = v18
                case 128:
                    switch fi16 {
                    case 9223372036854775807:
                        if false {
                            v1 = v1
                            fmt.Println(v2, v1, fi16)
                        }
                        fmt.Println(fi17)
                        var v19 string = "foo bar"
                        v1 -= 3.14159
                        var v20 int = -2147483648
                        var v21 uint = 571
                        _ = v19
                        _ = v20
                        _ = v21
                    case 995:
                        v22 := []int{(^(fi17)), (fi16 * fi17), fi17}
                        _ = v22
                    default:
                        delete(v3, 1)
                    }
                default:
                    v23 := GenS4{F0: "a"}
                    _ = v23
                }
                _ = fi17
            }
            _ = fi16
        }
        v24 := v3[1]
        if (42 >= 42) {
            if (!((1e10 * 0.1) >= (float32(((v1 + (v1 + (1e10 + (v1 * 0.1)))) * v1)) * (3.14159 * (1e10 - (0.1 / (float32((v1 * (((v1 - -2.25) + v1) + v1))) / 0.1))))))) {
                var v25 int32 = -1
                _ = v25
            }
            var v26 string = (("\n\t" + ("foo bar" + v24)) + "hello")
            v24 = v24
            _ = v26
        } else {
            v3[2147483648] = (((v24 + "\n\t") + "foo bar") + v24)
            v27 := GenS2{F0: 233, F1: false, F2: false}
            _ = v27
        }
        _ = v24
    }
    for fi28 := 0; fi28 < 6; fi28++ {
        if (0 >= 920) {
            v29 := []int{(fi28 & fi28), (fi28 | 4294967295)}
            v2 -= (v2 << ((18446744073709551615) & 63))
            for fi30 := 0; fi30 < 15; fi30++ {
                v1 = 1.5
                _ = fi30
            }
            _ = v29
        }
        var v31 uint = 18446744073709551615
        _ = fi28
        _ = v31
    }
    switch 42 {
    case 518:
        if ((^(v2)) > -1) {
            var v32 int = -77
            var v33 int16 = 975
            v34 := GenS5{F0: v32, F1: 1, F2: v2}
            _ = v32
            _ = v33
            _ = v34
        } else {
            var v35 bool = true
            var v36 int64 = (v2 * v2)
            _ = v35
            _ = v36
        }
    default:
        fmt.Println(v2, v2)
    }
    switch "\n\t" {
    case "":
        for fi37 := 0; fi37 < 16; fi37++ {
            fmt.Println(v1, v2)
            for fi38 := 0; fi38 < 9; fi38++ {
                if (!true) {
                    var v39 float32 = 0.1
                    _ = v39
                }
                if (!true) {
                    switch (fi38 | -457) {
                    case 87:
                        v40 := GenS3{F0: (fi38 & (fi37 * (fi37 + (fi37 & fi37)))), F1: "", F2: "foo bar"}
                        _ = v40
                    case 42:
                        v41 := GenS4{F0: ((("\n\t" + ("foo bar" + "\n\t")) + ("\n\t" + "foo bar")) + "hello")}
                        _ = v41
                    default:
                        v42 := GenS6{F0: 0, F1: 519}
                        _ = v42
                    }
                    var v43 uint8 = 23
                    fmt.Println(v2)
                    _ = v43
                }
                _ = fi38
            }
            v2 += (v2 << ((65535) & 63))
            _ = fi37
        }
    case "hello":
        v44 := GenS1{F0: "a", F1: 1, F2: ((("" + ("a" + (("\n\t" + "a") + "foo bar"))) + "\n\t") + ("\n\t" + "hello"))}
        _ = v44
    default:
        fmt.Println(v1, v2)
    }
    delete(v3, 65535)
    if (4294967295 <= 849) {
        fmt.Println(v2, v1)
    }
    v45 := []int{890, 9223372036854775807, 89, 191}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v45:", len(v45), v45)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v45
}
