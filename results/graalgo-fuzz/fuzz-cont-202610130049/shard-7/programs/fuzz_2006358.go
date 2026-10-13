package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 int32 }
func genFn0(p0 int, p1 string) uint32 {
    p1 = ((p1 + p1) + p1)
    fmt.Println(p1, p0)
    for fi6 := 0; fi6 < 1; fi6++ {
        p0 = p0
        _ = fi6
    }
    _ = p0
    _ = p1
    return 2
}

func main() {
    var v1 uint = 4294967295
    v2 := []uint32{65535, 610, 859, 2147483647}
    var v3 bool = false
    var v4 string = "hello"
    var v5 uint16 = 1
    var v7 uint64 = 4294967295
    fmt.Println(v3, v7)
    var v8 uint32 = 720
    v9 := GenS1{F0: false, F1: v1}
    if v3 {
        if (-2147483648 >= 300) {
            switch -571 {
            case 128:
                v10 := GenS2{F0: 42}
                _ = v10
            case 9223372036854775807:
                var v11 uint16 = 367
                _ = v11
            case 32767:
                v8 -= genFn0(-820, ("a" + ("hello" + ((((("hello" + v4) + ((v4 + "") + v4)) + ((v4 + "") + v4)) + v4) + v4))))
            }
            switch "a" {
            case "hello":
                var v12 uint32 = 2147483648
                _ = v12
            }
            if ((v5 << ((517) & 63)) == 40) {
                if v3 {
                    for fi13 := 0; fi13 < 1; fi13++ {
                        for fi14 := 0; fi14 < 20; fi14++ {
                            fmt.Println(v8, fi13, fi14)
                            _ = fi14
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = fi13
                    }
                }
                fmt.Println(v7, v1)
                if v3 {
                    v5 = uint16(v7)
                    v15 := GenS1{F0: (!true), F1: 18446744073709551615}
                    switch -522 {
                    case -295:
                        v5 += 0
                    case 651:
                        if (!(v5 < (v5 << (((v1 - v1)) & 63)))) {
                            if v3 {
                                var v16 uint8 = uint8(v8)
                                _ = v16
                            }
                            fmt.Println(v4, v7)
                            v8 = genFn0(int(v8), "foo bar")
                        }
                    }
                    _ = v15
                }
            } else {
                v17 := GenS1{F0: v3, F1: uint(v8)}
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v5 += (v5 % (((v5 | 32767)) | 1))
                _ = v17
            }
        }
        if len(v2) > 0 {
            v2[0] = genFn0(-1, v4)
        }
        var v18 int16 = 1
        _ = v18
    }
    v4 = "\n\t"
    fmt.Println(v7)
    if (-163 != -210) {
        var v19 int = -2147483648
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
