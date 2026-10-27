package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16 }
func genFn0(p0 uint8) string {
    v7 := []int{127, 42, -4}
    var v8 bool = true
    _ = p0
    _ = v7
    _ = v8
    return "foo bar"
}
func genFn1(p0 uint16) int8 {
    var v9 uint8 = 1
    fmt.Println(v9, p0)
    p0 -= p0
    _ = p0
    _ = v9
    return 12
}

func main() {
    var v1 uint = 18446744073709551615
    var v2 string = "\n\t"
    v3 := []int{-1}
    var v4 uint16 = 386
    v5 := []int{4294967295, 32767}
    var v6 uint32 = 447
    for fi10 := 0; fi10 < 1; fi10++ {
        _ = len(v3)
        if len(v3) > 0 {
            v3[0] = (fi10 - 195)
        }
        v11 := GenS1{F0: 42}
        _ = fi10
        _ = v11
    }
    var v12 uint8 = 219
    switch -2147483648 {
    case -1:
        var v13 uint8 = v12
        _ = v13
    }
    switch "" {
    case "a":
        if true {
            var v14 uint8 = 1
            var v15 uint64 = 18446744073709551615
            _ = v14
            _ = v15
        } else {
            if true {
                var v16 string = v2
                _ = v16
            } else {
                switch 913 {
                case 9223372036854775807:
                    if false {
                        switch -9223372036854775808 {
                        case 9223372036854775807:
                            if (!true) {
                                if len(v3) > 0 {
                                    v3[0] = -2147483648
                                }
                            }
                        case 128:
                            var v17 uint8 = v12
                            _ = v17
                        }
                        var v18 int16 = 0
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = v18
                    }
                case 667:
                    fmt.Println(v6)
                case 127:
                    for fi19 := 0; fi19 < 7; fi19++ {
                        v12 = (v12 % (((^(v12))) | 1))
                        _ = fi19
                    }
                default:
                    fmt.Println(v6)
                }
                if ((v2 >= ("" + (v2 + (("hello" + (("hello" + (v2 + v2)) + (v2 + (v2 + v2)))) + v2)))) && (false || true)) {
                    v2 = v2
                    var v20 int8 = 0
                    v21 := GenS1{F0: (^(v4))}
                    _ = v20
                    _ = v21
                }
            }
        }
    }
    v22 := GenS1{F0: 329}
    fmt.Println(v12, v1, v2)
    fmt.Println(v1)
    v4 = v4
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if false {
        var v23 uint64 = 0
        var v24 float64 = ((1.5 / (float64(1e10) * (1.5 - ((1.5 - -2.25) * 1e10)))) + 0.1)
        _ = v23
        _ = v24
    } else {
        if ((true && false) || (-32768 == int16(v6))) {
            var v25 string = ((("a" + v2) + "") + "hello")
            _ = v25
        } else {
            for fi26 := 0; fi26 < 7; fi26++ {
                var v27 int32 = 948
                _ = fi26
                _ = v27
            }
            if len(v3) > 0 {
                v3[0] = 32768
            }
            var v28 int64 = 149
            _ = v28
        }
    }
    v3 = append(v3, -914)
    var v29 int = -75
    if (!false) {
        if true {
            if len(v5) > 0 {
                v5[0] = (v29 + v29)
            }
            var v30 float32 = -2.25
            for fi31 := 0; fi31 < 6; fi31++ {
                for fi32 := 0; fi32 < 8; fi32++ {
                    if (42 <= 127) {
                        if (!(1 < 950)) {
                            var v33 uint16 = (v4 - v4)
                            if false {
                                v33 = (v4 - v4)
                            } else {
                                v34 := GenS2{F0: 0}
                                _ = v34
                            }
                            _ = v33
                        }
                    }
                    v12 = uint8(v4)
                    fmt.Println(v6)
                    _ = fi32
                }
                _ = fi31
            }
            _ = v30
        }
    }
    var v35 uint16 = (v4 << (((v1 ^ 0)) & 63))
    for fi36 := 0; fi36 < 8; fi36++ {
        switch (-(v29)) {
        case -587:
            switch (v29 + (fi36 * fi36)) {
            case 9223372036854775807:
                var v37 uint8 = uint8(v6)
                _ = v37
            }
        case 42:
            var v38 uint16 = (v35 << (((v1 + v1)) & 63))
            _ = v38
        case 418:
            var v39 int32 = 1
            _ = v39
        }
        var v40 int64 = 1
        _ = fi36
        _ = v40
    }
    fmt.Println(v29, v1, v29)
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v22
    _ = v29
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
