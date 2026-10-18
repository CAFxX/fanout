package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64; F1 uint; F2 string }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 string; F1 int16 }
type GenS6 struct { F0 string; F1 int32 }
type GenS7 struct { F0 uint8; F1 string; F2 uint16 }
func genFn0(p0 int, p1 int32, p2 uint32) string {
    var v4 int16 = 1000
    v5 := GenS1{F0: v4}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return (("a" + "a") + (("hello" + ("\n\t" + ("foo bar" + ("" + "")))) + "foo bar"))
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []uint32{42, 42, 0, 2147483648}
    var v3 string = "\n\t"
    v6 := GenS2{F0: 2}
    v1[462] = 4294967295
    v7 := GenS3{F0: -9223372036854775808, F1: 238, F2: v3}
    v3 = "hello"
    if (1.5 == 1e10) {
        if (!false) {
            v3 = ("\n\t" + "hello")
            var v8 uint32 = 65535
            var v9 uint = 733
            _ = v8
            _ = v9
        }
        fmt.Println(v3)
    } else {
        switch 32767 {
        case 2147483647:
            v3 = v3
        }
    }
    if false {
        for fi10 := 0; fi10 < 15; fi10++ {
            if (false && (0.0 == (0.1 - 0.0))) {
                for fi11 := 0; fi11 < 5; fi11++ {
                    var v12 float32 = 1.5
                    if (true && false) {
                        v3 = ""
                        var v13 int32 = 759
                        var v14 int64 = 0
                        _ = v13
                        _ = v14
                    }
                    _ = fi11
                    _ = v12
                }
                delete(v1, 2147483648)
            } else {
                v3 = ""
                var v15 int = fi10
                var v16 uint16 = 0
                _ = v15
                _ = v16
            }
            if true {
                var v17 float64 = ((((0.0 / 3.14159) - 3.14159) - 0.0) * 1.5)
                v18 := GenS3{F0: -9223372036854775808, F1: 42, F2: "foo bar"}
                _ = v17
                _ = v18
            } else {
                if (42 >= 2147483648) {
                    v19 := GenS3{F0: 1, F1: 0, F2: v3}
                    var v20 uint32 = 2147483648
                    var v21 int8 = 52
                    _ = v19
                    _ = v20
                    _ = v21
                } else {
                    var v22 bool = false
                    _ = v22
                }
            }
            _ = fi10
        }
        v3 = v3
        v23 := GenS4{F0: 1}
        _ = v23
    }
    if true {
        v24 := GenS2{F0: 4294967295}
        _ = v24
    }
    if (!false) {
        if ((-1 > -17) && true) {
            if (false && true) {
                v3 = ("hello" + v3)
                var v25 int8 = -1
                _ = v25
            }
            if ((!(!true)) || (!true)) {
                var v26 int8 = -128
                _ = len(v2)
                v26 -= (v26 >> ((0) & 63))
                _ = v26
            }
            _ = len(v2)
        } else {
            var v27 string = v3
            _ = v27
        }
        if false {
            for fi28 := 0; fi28 < 11; fi28++ {
                v3 = v3
                for fi29 := 0; fi29 < 20; fi29++ {
                    fmt.Println(fi29)
                    if (0.1 == (((((float64(0.0) - -2.25) * 0.1) - -2.25) * float64((0.1 / ((((-2.25 * (0.0 + -2.25)) * (1.5 - 1.5)) + 1.5) + float32((3.14159 - ((1e10 - ((float64(0.1) - float64(1e10)) + (-2.25 / 1.0))) + -2.25))))))) / -2.25)) {
                        var v30 int64 = -9223372036854775808
                        v31 := GenS5{F0: "foo bar", F1: -283}
                        _ = v30
                        _ = v31
                    }
                    _ = fi29
                }
                _ = fi28
            }
            switch 128 {
            case -2147483648:
                if false {
                    if false {
                        var v32 int32 = 1
                        _ = v32
                    } else {
                        v3 = v3
                        if (1e10 < 1.5) {
                            v3 = v3
                        }
                        var v33 bool = true
                        _ = v33
                    }
                    var v34 int16 = -968
                    _ = v34
                }
            case 127:
                if (false && ((((((3.14159 / 1.0) * 0.0) == 1e10) && (!(127 <= -117))) && true) && true)) {
                    v2 = append(v2, 2147483648)
                } else {
                    if (true && (-2147483648 >= -1)) {
                        v35 := GenS4{F0: 42}
                        _ = v35
                    }
                    var v36 int16 = 0
                    var v37 uint8 = 0
                    _ = v36
                    _ = v37
                }
            }
        }
        if true {
            v38 := GenS6{F0: v3, F1: 1}
            v39 := GenS6{F0: (v3 + ("" + v3)), F1: 933}
            v3 = ("a" + ((((v3 + v3) + (v3 + v3)) + (v3 + "")) + (v3 + v3)))
            _ = v38
            _ = v39
        }
    }
    v40 := GenS7{F0: 127, F1: "\n\t", F2: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v40
    _ = v6
    _ = v7
}
