package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8 }
func genFn0(p0 int32) int {
    if true {
        if (false && true) {
            if (!(!true)) {
                switch 0 {
                case 42:
                    var v4 float32 = float32(float64(1.5))
                    _ = v4
                case -667:
                    v5 := GenS1{F0: 301}
                    _ = v5
                case 4294967295:
                    if (!false) {
                        if false {
                            p0 = (p0 >> ((42) & 63))
                        } else {
                            p0 += 1
                            p0 += 2147483647
                            fmt.Println(p0)
                        }
                    }
                default:
                    var v6 float32 = ((1.5 - -2.25) - 3.14159)
                    _ = v6
                }
                p0 = p0
            } else {
                p0 = (p0 - p0)
                v7 := []int{32768}
                for fi8 := 0; fi8 < 18; fi8++ {
                    p0 = (^(p0))
                    p0 = p0
                    _ = fi8
                }
                _ = v7
            }
            fmt.Println(p0)
        } else {
            var v9 int64 = -9223372036854775808
            switch -200 {
            case 42:
                fmt.Println(p0, v9)
            default:
                if (!(889 > 886)) {
                    v10 := GenS2{F0: 42}
                    if false {
                        switch 153 {
                        case -822:
                            var v11 int64 = (-(v9))
                            _ = v11
                        case 42:
                            var v12 uint8 = 127
                            _ = v12
                        }
                        v13 := []int{32768, 127}
                        p0 -= p0
                        _ = v13
                    } else {
                        v14 := GenS3{F0: 42}
                        _ = v14
                    }
                    _ = v10
                } else {
                    v15 := GenS2{F0: -1}
                    _ = v15
                }
            }
            _ = v9
        }
    } else {
        v16 := []int{32768, 0}
        p0 += p0
        _ = v16
    }
    var v17 int = 4294967295
    var v18 int = (v17 >> ((42) & 63))
    fmt.Println(v17)
    _ = p0
    _ = v17
    _ = v18
    return (-(v18))
}
func genFn1() uint8 {
    if (0 > 2147483647) {
        for fi19 := 0; fi19 < 20; fi19++ {
            for fi20 := 0; fi20 < 2; fi20++ {
                v21 := GenS2{F0: -75}
                _ = fi20
                _ = v21
            }
            if ((42 >= 255) || ((-2147483648 < 74) && false)) {
                var v22 string = "\n\t"
                v23 := []int{fi19, (fi19 << ((18446744073709551615) & 63)), (fi19 * (fi19 / ((genFn0(995)) | 1)))}
                for fi24 := 0; fi24 < 13; fi24++ {
                    v22 = (((v22 + "\n\t") + v22) + "")
                    _ = len(v23)
                    var v25 float64 = ((3.14159 * (-2.25 - 0.1)) / 1.0)
                    _ = fi24
                    _ = v25
                }
                _ = v22
                _ = v23
            }
            var v26 int = (fi19 + (fi19 * 128))
            _ = fi19
            _ = v26
        }
    }
    v27 := []int{}
    var v28 int8 = 127
    var v29 int = 127
    switch (("foo bar" + (("\n\t" + (("" + "") + ("hello" + ""))) + "foo bar")) + "hello") {
    case "":
        v29 += (^(v29))
    case "a":
        fmt.Println(v28, v28)
    default:
        fmt.Println(v29, v29)
    }
    _ = v27
    _ = v28
    _ = v29
    return 11
}

func main() {
    var v1 string = "\n\t"
    var v2 int = -1
    var v3 int64 = 42
    var v30 float32 = 1e10
    v2 -= (^(v2))
    v31 := GenS3{F0: -128}
    fmt.Println(v3)
    var v32 uint64 = 713
    v33 := []int{1, v2, genFn0(163), v2}
    v33 = append(v33, 32768)
    v3 = v3
    var v34 int = genFn0(-2147483648)
    v30 = (v30 + 1e10)
    v32 = v32
    var v35 uint8 = 53
    v35 -= 42
    var v36 string = v1
    var v37 int64 = v3
    var v38 string = (v1 + v1)
    fmt.Println(v38)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
}
