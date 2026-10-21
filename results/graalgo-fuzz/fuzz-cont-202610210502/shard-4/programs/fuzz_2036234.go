package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0() string {
    var v6 uint8 = 2
    v6 = v6
    var v7 uint64 = 168
    v7 -= v7
    _ = v6
    _ = v7
    return "foo bar"
}

func main() {
    var v1 uint32 = 761
    var v2 uint32 = (v1 * (v1 & (v1 + v1)))
    var v3 uint8 = 149
    var v4 int16 = 42
    v5 := []int{-218, 9223372036854775807}
    switch "hello" {
    case "\n\t":
        if len(v5) > 0 {
            v5[0] = -923
        }
    case "foo bar":
        for fi8 := 0; fi8 < 16; fi8++ {
            v5 = append(v5, fi8)
            _ = fi8
        }
    case "hello":
        var v9 uint16 = 421
        _ = v9
    default:
        if false {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            switch 16 {
            case 0:
                if (!(70 != 2)) {
                    var v10 string = "hello"
                    _ = v10
                }
            case -1:
                if true {
                    for fi11 := 0; fi11 < 10; fi11++ {
                        v3 = 2
                        var v12 string = ("" + "\n\t")
                        _ = fi11
                        _ = v12
                    }
                    var v13 bool = true
                    var v14 float32 = -2.25
                    _ = v13
                    _ = v14
                } else {
                    v15 := GenS1{F0: 32768}
                    _ = v15
                }
            }
        }
    }
    if false {
        v2 -= v1
    } else {
        var v16 uint16 = 32767
        _ = v16
    }
    v17 := GenS1{F0: -2147483648}
    var v18 int16 = v4
    var v19 int16 = -32768
    var v20 int = -1
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v21 uint32 = (^(v2))
    var v22 int = v20
    if ((true || (61 < -36)) || (0.1 <= 0.0)) {
        switch v22 {
        case 4294967295:
            var v23 uint32 = (^(v2))
            _ = v23
        }
        fmt.Println(v1, v3, v22)
    } else {
        if true {
            switch "" {
            case "a":
                var v24 int = v22
                _ = v24
            default:
                if (false && (true || (("hello" < "\n\t") || false))) {
                    var v25 string = "foo bar"
                    _ = v25
                } else {
                    if (v3 > v3) {
                        if false {
                            var v26 uint32 = 474
                            var v27 uint64 = 2
                            _ = v26
                            _ = v27
                        }
                        for fi28 := 0; fi28 < 15; fi28++ {
                            v20 += (v22 * v22)
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            var v29 uint32 = v1
                            _ = fi28
                            _ = v29
                        }
                        v4 = v4
                    } else {
                        fmt.Println(v2, v3, v2)
                        for fi30 := 0; fi30 < 1; fi30++ {
                            var v31 int = (fi30 & (-(v20)))
                            v5 = append(v5, (v22 << ((217) & 63)))
                            _ = fi30
                            _ = v31
                        }
                    }
                }
            }
        } else {
            var v32 uint16 = 972
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            if true {
                var v33 int8 = 42
                _ = v33
            }
            _ = v32
        }
        v34 := GenS1{F0: v20}
        _ = v34
    }
    switch ((((("foo bar" + "hello") + "a") + "hello") + "\n\t") + ("hello" + "")) {
    case "\n\t":
        var v35 string = (("a" + "a") + ("a" + "foo bar"))
        _ = v35
    default:
        for fi36 := 0; fi36 < 5; fi36++ {
            var v37 int64 = 42
            var v38 float64 = float64(float32(1.5))
            _ = fi36
            _ = v37
            _ = v38
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
