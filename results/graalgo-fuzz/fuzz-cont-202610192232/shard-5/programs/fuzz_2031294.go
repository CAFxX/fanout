package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int8; F1 uint64 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint16 }

func main() {
    v1 := make(map[uint32]string)
    v2 := []uint32{}
    v3 := []uint32{669, 905, 476, 4294967295}
    if false {
        if (("foo bar" + "foo bar") >= "hello") {
            if (1 <= 2147483648) {
                var v4 uint64 = 2
                v4 = 1
                if true {
                    v4 += v4
                    v4 = v4
                }
                _ = v4
            } else {
                var v5 uint32 = 2
                v6 := GenS1{F0: ("" + ("a" + "hello")), F1: 42, F2: true}
                _ = v5
                _ = v6
            }
            var v7 int = 4294967295
            v8 := GenS1{F0: "foo bar", F1: 985, F2: true}
            _ = v7
            _ = v8
        } else {
            var v9 int16 = -1
            _ = v9
        }
        var v10 float32 = 3.14159
        _ = v10
    } else {
        for fi11 := 0; fi11 < 7; fi11++ {
            fmt.Println(fi11)
            v12 := GenS2{F0: 236}
            if false {
                for fi13 := 0; fi13 < 7; fi13++ {
                    if (!(!((42 > -798) || false))) {
                        fmt.Println(fi11, fi11)
                        var v14 float64 = float64(0.0)
                        _ = v14
                    } else {
                        var v15 uint = 818
                        _ = v15
                    }
                    v3 = append(v3, 65535)
                    fmt.Println(fi13, fi13)
                    _ = fi13
                }
                switch ((("\n\t" + ("a" + "a")) + ("foo bar" + "\n\t")) + "\n\t") {
                case "hello":
                    var v16 int64 = 558
                    _ = v16
                case "foo bar":
                    var v17 bool = false
                    _ = v17
                }
            } else {
                var v18 float32 = (((1.5 + 0.0) + 1.5) * -2.25)
                _ = v18
            }
            _ = fi11
            _ = v12
        }
        for fi19 := 0; fi19 < 20; fi19++ {
            v3 = append(v3, 1)
            var v20 uint16 = 210
            var v21 float64 = 0.1
            _ = fi19
            _ = v20
            _ = v21
        }
    }
    v3 = append(v3, 2147483647)
    v22 := GenS3{F0: 42}
    for fi23 := 0; fi23 < 19; fi23++ {
        _ = len(v1)
        switch -2147483648 {
        case -9223372036854775808:
            v2 = append(v2, 2147483647)
        case 128:
            if true {
                var v24 int32 = -132
                _ = v24
            }
        case 79:
            var v25 float32 = (float32(-2.25) - -2.25)
            _ = v25
        }
        var v26 int16 = 392
        _ = fi23
        _ = v26
    }
    var v27 bool = true
    var v28 uint8 = 128
    if v27 {
        if true {
            fmt.Println(v28, v28)
            v27 = (2147483647 <= 998)
            _ = len(v1)
        } else {
            if (1e10 >= 0.1) {
                fmt.Println(v28, v27)
            }
            v29 := GenS4{F0: 1, F1: 481}
            _ = v29
        }
        if (true && (-574 != 417)) {
            if ((!(!v27)) && (-2147483648 != -284)) {
                if (true && ((0.0 == ((1e10 * (0.1 - (float64(((0.1 - 1.5) + 0.0)) + float64((0.0 * ((1e10 * 0.1) + 0.1)))))) / ((1.5 / 0.1) - 3.14159))) && (!v27))) {
                    var v30 int8 = 100
                    _ = v30
                } else {
                    v27 = false
                    var v31 int8 = 100
                    var v32 uint8 = (v28 | 128)
                    _ = v31
                    _ = v32
                }
                var v33 uint64 = 613
                _ = v33
            }
            _ = len(v1)
            var v34 bool = (!(!true))
            _ = v34
        } else {
            var v35 int64 = -1
            if (float64((((0.1 * 1.5) - 0.1) * (-2.25 + 1.5))) != 0.1) {
                for fi36 := 0; fi36 < 6; fi36++ {
                    for fi37 := 0; fi37 < 5; fi37++ {
                        v27 = v27
                        v38 := GenS2{F0: v28}
                        var v39 int16 = 32767
                        _ = fi37
                        _ = v38
                        _ = v39
                    }
                    _ = fi36
                }
            }
            _ = v35
        }
        for fi40 := 0; fi40 < 12; fi40++ {
            fmt.Println(fi40, fi40)
            v28 = (v28 % ((v28) | 1))
            _ = fi40
        }
    }
    v27 = v27
    for fi41 := 0; fi41 < 20; fi41++ {
        fmt.Println(fi41, v27, v28)
        v42 := GenS5{F0: (v28 & v28)}
        var v43 int64 = 42
        _ = fi41
        _ = v42
        _ = v43
    }
    v27 = v27
    if len(v2) > 0 {
        v2[0] = 472
    }
    v44 := GenS6{F0: 401}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v22
    _ = v27
    _ = v28
    _ = v3
    _ = v44
}
