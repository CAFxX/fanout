package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 int64 }
type GenS2 struct { F0 bool; F1 int8; F2 uint8 }
type GenS3 struct { F0 int32; F1 int; F2 int16 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int = 4294967295
    v3 := make(map[int]int)
    v4 := make(map[int]uint32)
    v5 := make(map[string]uint32)
    if (!(!true)) {
        if false {
            if false {
                var v6 uint = 224
                _ = v6
            } else {
                if true {
                    var v7 int64 = -9223372036854775808
                    _ = v7
                } else {
                    var v8 int16 = 1
                    var v9 uint32 = 1
                    switch ("foo bar" + (("hello" + (("foo bar" + "\n\t") + "")) + "")) {
                    case "":
                        _ = len(v5)
                    }
                    _ = v8
                    _ = v9
                }
                for fi10 := 0; fi10 < 17; fi10++ {
                    v2 = (^(v2))
                    for fi11 := 0; fi11 < 11; fi11++ {
                        v12 := GenS1{F0: 255, F1: 186, F2: 9223372036854775807}
                        if (!(!(9223372036854775807 <= 42))) {
                            _ = len(v1)
                            var v13 int32 = 42
                            v2 = (fi10 << ((42) & 63))
                            _ = v13
                        }
                        _ = fi11
                        _ = v12
                    }
                    _ = fi10
                }
            }
        } else {
            var v14 int16 = 0
            v15 := []int{v2, 590}
            v14 += 1
            _ = v14
            _ = v15
        }
        for fi16 := 0; fi16 < 16; fi16++ {
            v2 = fi16
            if true {
                var v17 float64 = 1e10
                var v18 int16 = 32767
                _ = v17
                _ = v18
            }
            _ = fi16
        }
        v19 := GenS1{F0: 126, F1: 4294967295, F2: 903}
        _ = v19
    }
    switch 127 {
    case 0:
        fmt.Println(v2)
    case -1:
        for fi20 := 0; fi20 < 7; fi20++ {
            delete(v1, 65535)
            v21 := GenS1{F0: 127, F1: 18446744073709551615, F2: -1}
            if (!((float32(3.14159) + 0.1) <= float32((3.14159 - float64((1e10 - 3.14159)))))) {
                v22 := v4[v2]
                delete(v4, (-(fi20)))
                v22 = (v22 * (v22 >> ((18446744073709551615) & 63)))
                _ = v22
            }
            _ = fi20
            _ = v21
        }
    }
    var v23 float64 = (0.0 + 0.0)
    v23 += (float64(0.0) + (1e10 / 1.0))
    if (!(679 != 173)) {
        var v24 string = "a"
        v24 = ""
        _ = v24
    }
    for fi25 := 0; fi25 < 20; fi25++ {
        for fi26 := 0; fi26 < 9; fi26++ {
            var v27 int16 = 42
            _ = fi26
            _ = v27
        }
        var v28 uint64 = 2
        fmt.Println(v23)
        _ = fi25
        _ = v28
    }
    v29 := v3[(v2 * (v2 + v2))]
    for fi30 := 0; fi30 < 20; fi30++ {
        v31 := v5["a"]
        _ = fi30
        _ = v31
    }
    var v32 uint = 42
    v29 += (v2 & -207)
    if (((v2 + 4294967295) != v29) && (42 <= 128)) {
        var v33 float32 = (0.0 / 1.0)
        if (("a" + (("\n\t" + ("foo bar" + "")) + (("foo bar" + "foo bar") + "foo bar"))) <= ((("hello" + (("hello" + ("" + "")) + "foo bar")) + "a") + "")) {
            fmt.Println(v33, v2)
            v34 := GenS1{F0: 127, F1: 662, F2: 972}
            var v35 bool = true
            _ = v34
            _ = v35
        } else {
            v29 += v29
            for fi36 := 0; fi36 < 16; fi36++ {
                if true {
                    v37 := []int{v29, v29, fi36}
                    var v38 int64 = -1
                    for fi39 := 0; fi39 < 14; fi39++ {
                        _ = len(v37)
                        _ = fi39
                    }
                    _ = v37
                    _ = v38
                }
                v40 := GenS2{F0: false, F1: 0, F2: 128}
                _ = fi36
                _ = v40
            }
            v41 := []int{}
            _ = v41
        }
        _ = v33
    }
    v29 += (-(v29))
    v2 = v29
    v23 = 1.5
    for fi42 := 0; fi42 < 2; fi42++ {
        if true {
            v43 := []int{v2, (v2 ^ fi42), -9223372036854775808}
            _ = v43
        }
        v44 := GenS3{F0: 59, F1: 128, F2: 1000}
        _ = fi42
        _ = v44
    }
    v2 += (v2 | (v29 / (((v2 % (((-(v29))) | 1))) | 1)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v23
    _ = v29
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
