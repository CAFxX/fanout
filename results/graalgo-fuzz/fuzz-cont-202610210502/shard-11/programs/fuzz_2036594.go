package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64; F1 uint16; F2 uint }
type GenS3 struct { F0 int64 }
func genFn0(p0 uint64, p1 uint8) uint16 {
    fmt.Println(p0, p0)
    p1 += (p1 % ((p1) | 1))
    for fi4 := 0; fi4 < 3; fi4++ {
        var v5 string = "foo bar"
        if (fi4 < (^(fi4))) {
            var v6 float32 = 3.14159
            _ = v6
        } else {
            v7 := GenS1{F0: -9223372036854775808}
            _ = v7
        }
        fmt.Println(v5, v5, p1)
        _ = fi4
        _ = v5
    }
    p0 += p0
    v8 := []int{-1, -318, 32767}
    _ = p0
    _ = p1
    _ = v8
    return 449
}

func main() {
    var v1 uint32 = 1
    var v2 uint16 = 807
    var v3 uint16 = (v2 * v2)
    v9 := []int{}
    for fi10 := 0; fi10 < 9; fi10++ {
        v1 += 0
        v11 := GenS2{F0: -490, F1: (^(v2)), F2: 175}
        for fi12 := 0; fi12 < 19; fi12++ {
            switch (((("" + ("foo bar" + (("a" + "hello") + ((("a" + "\n\t") + "") + ("foo bar" + "\n\t"))))) + "\n\t") + "foo bar") + "foo bar") {
            case "":
                if (!true) {
                    v1 = 330
                    fmt.Println(v3)
                    v2 = (^(v3))
                } else {
                    switch "hello" {
                    case "\n\t":
                        v13 := GenS3{F0: -9223372036854775808}
                        _ = v13
                    case "hello":
                        fmt.Println(v2, fi10)
                    case "":
                        switch (fi10 << ((18) & 63)) {
                        case 1:
                            var v14 float32 = 1e10
                            _ = v14
                        default:
                            v15 := GenS1{F0: 38}
                            _ = v15
                        }
                        if true {
                            switch "hello" {
                            case "\n\t":
                                if (!((42 != 4294967295) || (false && (127 < -95)))) {
                                    if (false && ((fi10 - fi10) > -2147483648)) {
                                        var v16 int64 = -9223372036854775808
                                        fmt.Println(fi12, v3)
                                        _ = v16
                                    } else {
                                        fmt.Println(v3, fi12, fi10)
                                        if (1 >= -43) {
                                            v3 = 32767
                                            fmt.Println(fi12, fi12)
                                            switch (fi12 % (((^(fi12))) | 1)) {
                                            case -2147483648:
                                                fmt.Println(fi12)
                                            case -1:
                                                v3 = v3
                                            case 32768:
                                                if false {
                                                    v17 := GenS3{F0: -666}
                                                    _ = v17
                                                }
                                            }
                                        }
                                    }
                                }
                            default:
                                v2 = v3
                            }
                            var v18 string = ("" + ("a" + ("hello" + "")))
                            if false {
                                v1 -= v1
                            }
                            _ = v18
                        } else {
                            fmt.Println(fi12)
                            var v19 int32 = 772
                            _ = v19
                        }
                        _ = len(v9)
                    default:
                        v3 -= 2
                    }
                    switch "a" {
                    case "\n\t":
                        var v20 int8 = 107
                        _ = v20
                    }
                }
            case "\n\t":
                switch ("\n\t" + (("foo bar" + "foo bar") + ("\n\t" + "foo bar"))) {
                case "\n\t":
                    fmt.Println(v2, v3)
                default:
                    v2 += (v3 << ((1) & 63))
                }
            }
            if len(v9) > 0 {
                v9[0] = fi10
            }
            _ = fi12
        }
        _ = fi10
        _ = v11
    }
    var v21 bool = false
    var v22 uint = 2
    switch 4294967295 {
    case -940:
        v1 = 601
    case 4294967295:
        fmt.Println(v21, v2)
    }
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    fmt.Println(v1)
    fmt.Println(v21, v22, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v9
}
