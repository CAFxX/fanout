package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32; F1 int8 }
func genFn0(p0 int16, p1 string, p2 int16) int64 {
    if false {
        v4 := []int{42, 2147483647, 32767, 32767}
        if (!(float32((1e10 + (0.1 - 1e10))) >= 0.1)) {
            if false {
                for fi5 := 0; fi5 < 18; fi5++ {
                    v6 := GenS1{F0: 128}
                    if (-9223372036854775808 == 42) {
                        for fi7 := 0; fi7 < 6; fi7++ {
                            v8 := GenS1{F0: fi7}
                            _ = fi7
                            _ = v8
                        }
                    } else {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                    fmt.Println(p2, fi5)
                    _ = fi5
                    _ = v6
                }
                switch "\n\t" {
                case "a":
                    if true {
                        if len(v4) > 0 {
                            v4[0] = -9223372036854775808
                        }
                    }
                case "\n\t":
                    switch 127 {
                    case -278:
                        fmt.Println(p0, p0, p2)
                    case 42:
                        fmt.Println(p0, p1)
                    }
                default:
                    fmt.Println(p0)
                }
                var v9 float64 = -2.25
                _ = v9
            } else {
                p1 = p1
                if false {
                    p1 = (p1 + "a")
                    if len(v4) > 0 {
                        v4[0] = 0
                    }
                    v10 := GenS1{F0: -209}
                    _ = v10
                }
            }
            if false {
                _ = len(v4)
            }
        }
        var v11 int32 = -2147483648
        _ = v11
        _ = v4
    } else {
        v12 := []int{}
        _ = v12
    }
    if (-9223372036854775808 < 4294967295) {
        fmt.Println(p2)
    } else {
        var v13 uint = 4294967295
        var v14 float32 = 1.5
        v14 = 0.1
        _ = v13
        _ = v14
    }
    p0 -= 32767
    if true {
        fmt.Println(p0, p0, p1)
        p2 -= 42
    } else {
        for fi15 := 0; fi15 < 12; fi15++ {
            var v16 float64 = -2.25
            _ = fi15
            _ = v16
        }
        p0 = (p0 % ((p0) | 1))
        for fi17 := 0; fi17 < 2; fi17++ {
            v18 := []int{(^(fi17)), fi17, fi17, -2147483648}
            _ = fi17
            _ = v18
        }
    }
    var v19 int = -9223372036854775808
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    return 1
}

func main() {
    var v1 uint = 18446744073709551615
    var v2 int8 = -112
    v3 := []string{"\n\t"}
    for fi20 := 0; fi20 < 7; fi20++ {
        v3 = append(v3, "a")
        _ = fi20
    }
    var v21 int32 = 2147483647
    if false {
        v21 = (v21 * (v21 + -2147483648))
    } else {
        v2 += 127
        v21 -= v21
    }
    v22 := GenS2{F0: 128}
    for fi23 := 0; fi23 < 6; fi23++ {
        for fi24 := 0; fi24 < 15; fi24++ {
            if len(v3) > 0 {
                v3[0] = "foo bar"
            }
            _ = fi24
        }
        _ = fi23
    }
    v3 = append(v3, ("hello" + "foo bar"))
    v3 = append(v3, "hello")
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch 9223372036854775807 {
    case 4294967295:
        for fi25 := 0; fi25 < 4; fi25++ {
            var v26 uint8 = 255
            _ = fi25
            _ = v26
        }
    default:
        v27 := GenS2{F0: 42}
        _ = v27
    }
    v28 := GenS1{F0: 257}
    var v29 string = "a"
    for fi30 := 0; fi30 < 5; fi30++ {
        var v31 string = "\n\t"
        var v32 int64 = -9223372036854775808
        if ((v32 & v32) != v32) {
            v29 = (v31 + (v29 + "foo bar"))
            switch ((v31 + ("foo bar" + "\n\t")) + "a") {
            case "hello":
                for fi33 := 0; fi33 < 2; fi33++ {
                    v34 := GenS3{F0: 2147483648, F1: v2}
                    _ = fi33
                    _ = v34
                }
            }
            if false {
                v1 -= v1
                switch "hello" {
                case "a":
                    var v35 uint16 = 65535
                    _ = v35
                }
                var v36 float64 = -2.25
                _ = v36
            }
        }
        _ = fi30
        _ = v31
        _ = v32
    }
    var v37 float32 = ((-2.25 + -2.25) * 1.5)
    v21 = (v21 << (((^(v1))) & 63))
    for fi38 := 0; fi38 < 16; fi38++ {
        var v39 int64 = 9223372036854775807
        var v40 float64 = 0.0
        _ = fi38
        _ = v39
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v28
    _ = v29
    _ = v3
    _ = v37
}
