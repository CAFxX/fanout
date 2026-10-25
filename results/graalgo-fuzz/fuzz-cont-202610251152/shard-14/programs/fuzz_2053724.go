package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 int; F1 int16; F2 int64 }

func main() {
    var v1 uint8 = 128
    v2 := make(map[string]uint32)
    var v3 bool = true
    var v4 uint8 = 91
    delete(v2, "a")
    v1 += 255
    var v5 uint16 = 0
    var v6 int = -2147483648
    if (!v3) {
        v7 := GenS1{F0: v4}
        fmt.Println(v4, v5)
        _ = v7
    } else {
        var v8 float64 = 1e10
        _ = v8
    }
    delete(v2, "\n\t")
    var v9 int64 = -9223372036854775808
    for fi10 := 0; fi10 < 15; fi10++ {
        v4 = v1
        for fi11 := 0; fi11 < 6; fi11++ {
            switch (fi11 % ((32768) | 1)) {
            case 32768:
                var v12 uint = 42
                _ = v12
            case 85:
                v13 := []int{}
                _ = v13
            case 1:
                v14 := []int{fi11, (v6 % ((4294967295) | 1)), fi11, (v6 + v6)}
                _ = v14
            default:
                if ((fi10 + (^(fi10))) <= 42) {
                    var v15 int32 = 2147483647
                    v1 -= v4
                    v16 := GenS1{F0: (v1 ^ (v1 ^ (^(v4))))}
                    _ = v15
                    _ = v16
                } else {
                    if (1000 > -32768) {
                        var v17 uint64 = 4294967295
                        v6 -= 127
                        v5 -= (v5 + (^(v5)))
                        _ = v17
                    } else {
                        var v18 int16 = -760
                        var v19 int32 = 2147483647
                        v5 += v5
                        _ = v18
                        _ = v19
                    }
                    v9 = (v9 - (v9 % (((v9 | -919)) | 1)))
                }
            }
            v20 := GenS1{F0: (v1 - (v1 >> ((838) & 63)))}
            switch fi10 {
            case -274:
                var v21 uint64 = 481
                _ = v21
            case -9223372036854775808:
                fmt.Println(fi11, v9, v9)
            }
            var v22 int32 = -343
            _ = fi11
            _ = v20
            _ = v22
        }
        fmt.Println(v6, v9, v6)
        _ = fi10
    }
    if (898 != 42) {
        for fi23 := 0; fi23 < 16; fi23++ {
            v24 := GenS2{F0: (("hello" + (("a" + "") + "foo bar")) + ""), F1: (0 != 0)}
            for fi25 := 0; fi25 < 3; fi25++ {
                v26 := GenS2{F0: ("hello" + ((("a" + "hello") + "a") + "hello")), F1: (!(v3 && true))}
                _ = fi25
                _ = v26
            }
            _ = fi23
            _ = v24
        }
        switch 409 {
        case 32768:
            fmt.Println(v5, v6)
        default:
            v27 := GenS1{F0: v4}
            _ = v27
        }
    }
    fmt.Println(v5)
    v6 += v6
    for fi28 := 0; fi28 < 14; fi28++ {
        switch "hello" {
        case "":
            if true {
                var v29 uint16 = (v5 - 42)
                for fi30 := 0; fi30 < 17; fi30++ {
                    fmt.Println(v5, v5, v5)
                    _ = fi30
                }
                v4 -= v1
                _ = v29
            } else {
                v3 = (1.5 > 0.1)
                for fi31 := 0; fi31 < 15; fi31++ {
                    var v32 int = -94
                    _ = fi31
                    _ = v32
                }
                v33 := []int{v6}
                _ = v33
            }
        default:
            v34 := GenS3{F0: 4294967295, F1: -32768, F2: v9}
            _ = v34
        }
        fmt.Println(v4)
        var v35 float32 = 1e10
        _ = fi28
        _ = v35
    }
    v36 := v2[("hello" + "a")]
    v4 += (v1 % (((v4 << ((2) & 63))) | 1))
    if (false && (!((v5 / (((^(v5))) | 1)) <= v5))) {
        v5 -= v5
        var v37 uint64 = 2
        _ = v37
    }
    var v38 uint64 = 768
    for fi39 := 0; fi39 < 9; fi39++ {
        switch ("hello" + "hello") {
        case "":
            var v40 int = (-(fi39))
            _ = v40
        default:
            var v41 uint8 = v4
            _ = v41
        }
        for fi42 := 0; fi42 < 1; fi42++ {
            v43 := v2[(("" + (("foo bar" + "hello") + ((((("foo bar" + "a") + ((("" + ("foo bar" + (("" + "foo bar") + (("foo bar" + ("a" + "a")) + "hello")))) + ("\n\t" + "")) + "")) + "\n\t") + "hello") + ("foo bar" + "hello")))) + ("a" + ("hello" + "a")))]
            v36 += (^(v36))
            var v44 int64 = 0
            _ = fi42
            _ = v43
            _ = v44
        }
        _ = fi39
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
