package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 uint8; F1 int64; F2 uint64 }
type GenS3 struct { F0 bool; F1 int16; F2 int }
type GenS4 struct { F0 uint }
func genFn0(p0 int32, p1 int16) int16 {
    for fi4 := 0; fi4 < 3; fi4++ {
        v5 := []int{370, (^(fi4))}
        for fi6 := 0; fi6 < 9; fi6++ {
            if true {
                var v7 uint8 = 255
                if (true || (41 == 0)) {
                    var v8 int32 = (p0 ^ -2147483648)
                    _ = v8
                } else {
                    v9 := GenS1{F0: "", F1: (fi4 ^ (fi4 ^ -795))}
                    _ = v9
                }
                _ = v7
            }
            _ = fi6
        }
        fmt.Println(p0, p1, fi4)
        _ = fi4
        _ = v5
    }
    for fi10 := 0; fi10 < 1; fi10++ {
        for fi11 := 0; fi11 < 13; fi11++ {
            p1 -= -1
            _ = fi11
        }
        switch (fi10 % ((fi10) | 1)) {
        case -246:
            if true {
                var v12 int = (fi10 << ((672) & 63))
                _ = v12
            }
        case -652:
            for fi13 := 0; fi13 < 9; fi13++ {
                v14 := []int{fi10, fi10}
                _ = fi13
                _ = v14
            }
        case -9223372036854775808:
            v15 := []int{1, fi10}
            _ = v15
        }
        _ = fi10
    }
    v16 := GenS1{F0: "foo bar", F1: 128}
    p1 = p1
    for fi17 := 0; fi17 < 1; fi17++ {
        for fi18 := 0; fi18 < 16; fi18++ {
            fmt.Println(p0, p0)
            _ = fi18
        }
        _ = fi17
    }
    _ = p0
    _ = p1
    _ = v16
    return (p1 << ((2) & 63))
}
func genFn1(p0 int32, p1 bool) string {
    var v19 int16 = -60
    v20 := []int{-1, 9223372036854775807, 69}
    _ = p0
    _ = p1
    _ = v19
    _ = v20
    return "a"
}

func main() {
    v1 := []string{"", "hello", (("a" + "hello") + ""), ("hello" + "hello")}
    var v2 int16 = 1
    var v3 int32 = 1
    v2 += 32767
    v2 -= (v2 << ((65535) & 63))
    var v21 uint16 = 178
    fmt.Println(v2, v2, v3)
    var v22 uint16 = (^(v21))
    fmt.Println(v3, v21)
    if true {
        var v23 bool = (!(v2 >= (^(v2))))
        var v24 uint16 = v21
        var v25 int64 = -9223372036854775808
        _ = v23
        _ = v24
        _ = v25
    }
    fmt.Println(v21, v21)
    var v26 uint32 = 4294967295
    if false {
        fmt.Println(v3, v2)
    } else {
        if true {
            v27 := GenS1{F0: "foo bar", F1: 32768}
            for fi28 := 0; fi28 < 13; fi28++ {
                v1 = append(v1, "")
                v29 := GenS2{F0: 42, F1: -726, F2: 668}
                if len(v1) > 0 {
                    v1[0] = "foo bar"
                }
                _ = fi28
                _ = v29
            }
            v2 -= v2
            _ = v27
        } else {
            v30 := GenS3{F0: true, F1: 1000, F2: -802}
            _ = v30
        }
        var v31 string = ("hello" + "hello")
        v32 := GenS3{F0: (!true), F1: -32768, F2: 2147483647}
        _ = v31
        _ = v32
    }
    if (1 < -1000) {
        v26 += (^(v26))
        if (1.5 <= float64(0.0)) {
            fmt.Println(v3, v26)
        } else {
            v22 = v22
            if false {
                v22 -= (^(v22))
                switch "" {
                case "\n\t":
                    if true {
                        var v33 uint8 = 2
                        fmt.Println(v33, v3, v21)
                        switch -9223372036854775808 {
                        case 0:
                            var v34 float64 = ((0.0 + 3.14159) - ((3.14159 + 3.14159) / float64((-2.25 + 0.1))))
                            _ = v34
                        default:
                            fmt.Println(v33)
                        }
                        _ = v33
                    }
                case "":
                    var v35 int8 = 42
                    _ = v35
                default:
                    v2 -= (-(v2))
                }
            } else {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v21 = 65535
            }
            fmt.Println(v21, v2)
        }
        for fi36 := 0; fi36 < 3; fi36++ {
            v1 = append(v1, (("a" + ("\n\t" + "hello")) + ((((("\n\t" + "foo bar") + "\n\t") + "a") + "foo bar") + ((("a" + "a") + "foo bar") + ("hello" + "a")))))
            v37 := GenS4{F0: 65535}
            _ = fi36
            _ = v37
        }
    } else {
        v38 := GenS2{F0: 42, F1: -205, F2: 794}
        fmt.Println(v26)
        _ = v38
    }
    if false {
        fmt.Println(v26, v2)
    }
    v26 = (v26 / (((v26 * v26)) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v26
    _ = v3
}
