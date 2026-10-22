package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 uint32 }

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 bool = false
    v3 := []string{"\n\t"}
    var v4 uint16 = 814
    var v5 int8 = 127
    if (!(667 == 638)) {
        v4 = 760
        if v2 {
            if ((-692 != -9223372036854775808) || (-2147483648 <= -918)) {
                v3 = append(v3, ("hello" + (("foo bar" + "hello") + ("a" + "a"))))
                v6 := GenS1{F0: 4294967295, F1: 9223372036854775807, F2: 4294967295}
                _ = v6
            }
        }
    } else {
        var v7 uint32 = 65535
        var v8 uint = 2
        var v9 uint64 = (v1 % ((v1) | 1))
        _ = v7
        _ = v8
        _ = v9
    }
    var v10 uint = 981
    if len(v3) > 0 {
        v3[0] = ""
    }
    v10 = 904
    v11 := GenS1{F0: 9223372036854775807, F1: -1, F2: uint32(v10)}
    var v12 float32 = (((float32((((float64(1e10) - 1e10) + 3.14159) / 0.1)) + ((1e10 + 0.1) - ((0.1 + float32(1.5)) / float32(1.5)))) * float32(0.0)) + (0.0 - 0.0))
    v5 += (v5 / ((v5) | 1))
    if false {
        v4 = 976
        fmt.Println(v10)
    }
    v5 += (^(v5))
    var v13 int8 = (-(v5))
    if (42 > v10) {
        var v14 uint64 = v1
        if len(v3) > 0 {
            v3[0] = "foo bar"
        }
        _ = v14
    }
    fmt.Println(v13, v2, v10)
    v10 = 622
    switch "\n\t" {
    case "foo bar":
        var v15 float64 = 3.14159
        _ = v15
    default:
        for fi16 := 0; fi16 < 15; fi16++ {
            if (true && false) {
                switch (fi16 | fi16) {
                case 2147483647:
                    switch "a" {
                    case "":
                        v12 += v12
                    case "\n\t":
                        for fi17 := 0; fi17 < 18; fi17++ {
                            var v18 uint16 = (v4 << ((42) & 63))
                            v13 = v13
                            var v19 uint = (v10 % ((v10) | 1))
                            _ = fi17
                            _ = v18
                            _ = v19
                        }
                    case "foo bar":
                        var v20 float64 = ((float64((0.0 / (3.14159 - v12))) / (float64(3.14159) - 1e10)) / (float64(((v12 + v12) * float32(0.1))) * float64((0.0 + (1e10 / float32((-2.25 + ((3.14159 / 0.1) * 0.0))))))))
                        _ = v20
                    default:
                        _ = len(v3)
                    }
                }
            } else {
                switch ("hello" + ("" + "foo bar")) {
                case "foo bar":
                    var v21 uint16 = (v4 ^ 860)
                    _ = v21
                case "a":
                    if false {
                        for fi22 := 0; fi22 < 13; fi22++ {
                            if v2 {
                                var v23 uint32 = 2
                                _ = v23
                            }
                            var v24 uint16 = v4
                            _ = fi22
                            _ = v24
                        }
                        v1 += (v1 % (((v1 ^ v1)) | 1))
                        if (!(!(v10 != v10))) {
                            v12 = v12
                            v10 += uint(v1)
                        } else {
                            var v25 int16 = 32767
                            switch ("a" + "a") {
                            case "hello":
                                var v26 string = ("\n\t" + "\n\t")
                                _ = v26
                            case "foo bar":
                                v4 -= v4
                            case "a":
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            default:
                                v10 -= (v10 | v10)
                            }
                            _ = v25
                        }
                    }
                }
                var v27 uint = v10
                _ = v27
            }
            var v28 uint32 = 4294967295
            _ = fi16
            _ = v28
        }
    }
    v10 -= 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
