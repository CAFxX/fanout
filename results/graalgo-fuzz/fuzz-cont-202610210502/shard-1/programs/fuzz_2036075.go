package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 int64 }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 string; F1 string }
type GenS4 struct { F0 int64; F1 uint32 }
type GenS5 struct { F0 int8; F1 int }
type GenS6 struct { F0 string; F1 uint64; F2 uint16 }
func genFn0() int16 {
    v6 := []int{32767}
    v7 := GenS1{F0: 472, F1: 526, F2: 42}
    _ = v6
    _ = v7
    return -1
}

func main() {
    var v1 int = -1
    var v2 int64 = 0
    var v3 int32 = -2147483648
    v4 := make(map[uint32]string)
    var v5 uint16 = 1
    v4[590] = "hello"
    v2 += 9223372036854775807
    v3 -= v3
    v1 = v1
    v1 = (v1 | v1)
    if false {
        v5 += (^(v5))
        v8 := GenS2{F0: 42, F1: v5}
        for fi9 := 0; fi9 < 8; fi9++ {
            switch 2147483647 {
            case 323:
                v10 := GenS3{F0: "foo bar", F1: ("\n\t" + ("a" + ""))}
                _ = v10
            case 4294967295:
                var v11 uint32 = 0
                _ = v11
            case 443:
                v2 = v2
            default:
                v12 := []int{fi9, -957, -9223372036854775808, v1}
                _ = v12
            }
            _ = fi9
        }
        _ = v8
    }
    v13 := GenS4{F0: v2, F1: 2147483648}
    var v14 int8 = -1
    v5 = (v5 & v5)
    if ("foo bar" >= (("hello" + ("hello" + "\n\t")) + "hello")) {
        var v15 uint64 = 2
        _ = v15
    } else {
        if false {
            v16 := []int{32767, -611, (v1 % (((v1 - (v1 - int(v5)))) | 1))}
            if true {
                fmt.Println(v3, v2, v14)
                if false {
                    v17 := v4[4294967295]
                    fmt.Println(v1)
                    _ = v17
                }
                if len(v16) >= 2 {
                    v16 = v16[1:len(v16)]
                }
            } else {
                v1 += v1
            }
            _ = v16
        } else {
            var v18 uint8 = 159
            _ = v18
        }
        for fi19 := 0; fi19 < 16; fi19++ {
            var v20 int64 = (^(v2))
            for fi21 := 0; fi21 < 6; fi21++ {
                if (false || (890 > 2147483648)) {
                    if true {
                        fmt.Println(v20, v5)
                    }
                    v20 -= 0
                    v5 -= 65535
                }
                _ = fi21
            }
            v22 := GenS5{F0: (v14 % (((v14 | v14)) | 1)), F1: fi19}
            _ = fi19
            _ = v20
            _ = v22
        }
    }
    for fi23 := 0; fi23 < 5; fi23++ {
        v14 = v14
        _ = fi23
    }
    if ("foo bar" >= "") {
        v4[4294967295] = "\n\t"
        v14 = (^(v14))
    } else {
        for fi24 := 0; fi24 < 14; fi24++ {
            var v25 int8 = v14
            var v26 string = "hello"
            switch (-(fi24)) {
            case 9223372036854775807:
                v27 := GenS5{F0: (v25 >> ((625) & 63)), F1: v1}
                _ = v27
            case 1:
                var v28 int16 = 141
                _ = v28
            }
            _ = fi24
            _ = v25
            _ = v26
        }
        v29 := GenS6{F0: "\n\t", F1: 332, F2: (^(v5))}
        var v30 int8 = v14
        _ = v29
        _ = v30
    }
    switch "" {
    case "foo bar":
        for fi31 := 0; fi31 < 19; fi31++ {
            var v32 int64 = v2
            v14 = (v14 * v14)
            var v33 uint32 = 65535
            _ = fi31
            _ = v32
            _ = v33
        }
    }
    v2 += (-(v2))
    if (65535 == 34) {
        v14 += (^(v14))
    } else {
        v4[2147483648] = ("a" + ((("" + "a") + "hello") + "a"))
        v14 += v14
        var v34 uint16 = (v5 | (v5 + (v5 & v5)))
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
