package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string; F1 uint }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint32; F1 bool; F2 uint64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int8; F1 uint8 }
func genFn0() int {
    var v5 float64 = float64(((((1.5 * float32(((float64(((0.1 * 0.1) / 1e10)) / 3.14159) * float64(3.14159)))) + (3.14159 - (float32(((float64(-2.25) * (1e10 / 1.0)) + 1.5)) * 3.14159))) * (float32(0.1) + (1.5 + 3.14159))) + 1e10))
    v6 := []int{}
    v5 = v5
    v7 := GenS1{F0: 0}
    v5 = ((v5 - float64((-2.25 / -2.25))) * (1e10 - 0.0))
    _ = v5
    _ = v6
    _ = v7
    return -400
}
func genFn1() uint {
    switch (("a" + "\n\t") + "foo bar") {
    case "a":
        v8 := GenS1{F0: 1}
        _ = v8
    }
    var v9 int64 = 0
    var v10 int32 = -827
    if (!true) {
        v11 := GenS1{F0: 0}
        var v12 uint64 = 642
        _ = v11
        _ = v12
    }
    _ = v10
    _ = v9
    return 84
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]bool)
    var v3 uint8 = 70
    var v4 uint32 = 4294967295
    v13 := GenS1{F0: 0}
    v4 = (v4 & v4)
    if true {
        var v14 float32 = (0.0 * float32(0.0))
        for fi15 := 0; fi15 < 19; fi15++ {
            for fi16 := 0; fi16 < 14; fi16++ {
                fmt.Println(fi15)
                _ = fi16
            }
            v17 := GenS2{F0: "foo bar", F1: 0}
            _ = fi15
            _ = v17
        }
        _ = v14
    }
    v18 := GenS3{F0: "hello"}
    switch "foo bar" {
    case "foo bar":
        v3 -= v3
    default:
        var v19 int = 931
        _ = v19
    }
    v1[v4] = (uint(v3) >= 1)
    v3 += v3
    if true {
        v20 := GenS4{F0: v4, F1: (32767 == -1), F2: 0}
        var v21 float64 = (1e10 - 3.14159)
        _ = v20
        _ = v21
    } else {
        fmt.Println(v4, v4)
        switch "hello" {
        case "":
            v3 -= 237
        case "hello":
            v22 := GenS2{F0: ("" + "a"), F1: 2}
            _ = v22
        default:
            fmt.Println(v3)
        }
    }
    if true {
        var v23 string = ""
        _ = len(v1)
        var v24 bool = (!true)
        _ = v23
        _ = v24
    } else {
        var v25 int64 = -9223372036854775808
        _ = v25
    }
    var v26 uint8 = (^(v3))
    fmt.Println(v4)
    for fi27 := 0; fi27 < 10; fi27++ {
        for fi28 := 0; fi28 < 1; fi28++ {
            if false {
                v29 := GenS3{F0: "hello"}
                var v30 uint64 = 672
                if false {
                    if false {
                        fmt.Println(v30, v3, v4)
                    } else {
                        var v31 string = ""
                        _ = v31
                    }
                    switch 176 {
                    case 2147483647:
                        _ = len(v2)
                    case 128:
                        v4 -= (v4 >> ((4294967295) & 63))
                    default:
                        v32 := []int{}
                        _ = v32
                    }
                    v33 := GenS3{F0: ""}
                    _ = v33
                } else {
                    var v34 float32 = float32((1.5 * 3.14159))
                    var v35 bool = (!(v4 >= (v4 + (^(v4)))))
                    switch "" {
                    case "foo bar":
                        v36 := GenS3{F0: (("hello" + "") + "\n\t")}
                        _ = v36
                    case "a":
                        if v35 {
                            var v37 uint64 = (v30 * (^(v30)))
                            _ = v37
                        } else {
                            if (!((v26 ^ (v3 - 3)) >= 0)) {
                                v38 := v1[(^(v4))]
                                var v39 uint64 = v30
                                var v40 uint16 = 427
                                _ = v38
                                _ = v39
                                _ = v40
                            } else {
                                var v41 uint64 = v30
                                _ = v41
                            }
                            v42 := GenS5{F0: fi28}
                            switch fi28 {
                            case 9223372036854775807:
                                v43 := GenS2{F0: "hello", F1: genFn1()}
                                _ = v43
                            }
                            _ = v42
                        }
                    }
                    _ = v34
                    _ = v35
                }
                _ = v29
                _ = v30
            } else {
                v44 := GenS6{F0: 127, F1: v26}
                var v45 int8 = 0
                _ = v44
                _ = v45
            }
            v46 := GenS6{F0: 118, F1: (^(v26))}
            _ = fi28
            _ = v46
        }
        switch (("foo bar" + "foo bar") + ("hello" + "hello")) {
        case "hello":
            var v47 uint16 = 255
            _ = v47
        case "foo bar":
            v26 -= (^(v3))
        }
        var v48 int64 = 355
        _ = fi27
        _ = v48
    }
    var v49 int8 = 127
    var v50 int16 = 32767
    var v51 string = (("foo bar" + "foo bar") + "a")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v13
    _ = v18
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v49
    _ = v50
    _ = v51
}
