package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 uint8, p1 uint16) string {
    if false {
        switch -9223372036854775808 {
        case 4294967295:
            if (-32768 > 0) {
                p1 = p1
            } else {
                v6 := GenS1{F0: 1}
                p0 = 211
                var v7 int16 = 1
                _ = v6
                _ = v7
            }
        default:
            if false {
                var v8 int = 128
                for fi9 := 0; fi9 < 1; fi9++ {
                    var v10 int64 = int64(v8)
                    _ = fi9
                    _ = v10
                }
                for fi11 := 0; fi11 < 17; fi11++ {
                    v12 := GenS1{F0: 0}
                    _ = fi11
                    _ = v12
                }
                _ = v8
            }
        }
        switch -9223372036854775808 {
        case -967:
            for fi13 := 0; fi13 < 19; fi13++ {
                var v14 int32 = -588
                if (9223372036854775807 < 488) {
                    var v15 int8 = -2
                    _ = v15
                } else {
                    switch "" {
                    case "hello":
                        if (-1 <= 42) {
                            p0 = (p0 << ((4294967295) & 63))
                            for fi16 := 0; fi16 < 2; fi16++ {
                                p1 += (^(p1))
                                p1 -= p1
                                _ = fi16
                            }
                        }
                    default:
                        var v17 float64 = -2.25
                        _ = v17
                    }
                    v18 := []int{fi13, (fi13 << ((1) & 63))}
                    _ = v18
                }
                _ = fi13
                _ = v14
            }
        default:
            var v19 int32 = 2147483647
            _ = v19
        }
    }
    var v20 int16 = 32767
    switch 127 {
    case 263:
        for fi21 := 0; fi21 < 10; fi21++ {
            var v22 int64 = 42
            _ = fi21
            _ = v22
        }
    }
    if true {
        v20 = (v20 - 0)
        fmt.Println(p0, v20)
        v20 += (v20 | (v20 >> ((42) & 63)))
    } else {
        v23 := []int{}
        if (-2147483648 == -443) {
            if len(v23) >= 2 {
                v23 = v23[1:len(v23)]
            }
            v24 := GenS2{F0: 0}
            _ = len(v23)
            _ = v24
        }
        _ = v23
    }
    _ = p0
    _ = p1
    _ = v20
    return "a"
}

func main() {
    var v1 string = "hello"
    var v2 float64 = 1e10
    var v3 uint64 = 420
    v4 := make(map[string]int)
    var v5 string = (v1 + "a")
    for fi25 := 0; fi25 < 3; fi25++ {
        if (0 == 912) {
            v3 = v3
            for fi26 := 0; fi26 < 1; fi26++ {
                var v27 int32 = 406
                v28 := []int{}
                if false {
                    switch (fi26 | -711) {
                    case 1:
                        v29 := GenS2{F0: v3}
                        _ = v29
                    }
                }
                _ = fi26
                _ = v27
                _ = v28
            }
        } else {
            fmt.Println(v5)
            var v30 string = (v1 + (v1 + "a"))
            switch fi25 {
            case -150:
                if true {
                    switch (v1 + v30) {
                    case "foo bar":
                        for fi31 := 0; fi31 < 19; fi31++ {
                            v32 := GenS1{F0: -32768}
                            _ = fi31
                            _ = v32
                        }
                    default:
                        var v33 uint64 = 0
                        _ = v33
                    }
                    var v34 int32 = -761
                    _ = v34
                }
            default:
                if (0 > 1) {
                    v35 := GenS3{F0: 128}
                    v2 += (v2 / float64(((0.0 - -2.25) - ((1e10 / ((-2.25 - float32(1.5)) / 1.0)) - float32((1.5 - v2))))))
                    _ = v35
                } else {
                    fmt.Println(v30, v5, v30)
                    v1 = v5
                }
            }
            _ = v30
        }
        _ = fi25
    }
    var v36 bool = (false || true)
    fmt.Println(v1, v3)
    for fi37 := 0; fi37 < 20; fi37++ {
        if (!(-2.25 >= (0.0 + 3.14159))) {
            var v38 int32 = 2147483647
            _ = v38
        }
        _ = fi37
    }
    var v39 int64 = -594
    var v40 int64 = -695
    v40 = v39
    v41 := v4[("foo bar" + v1)]
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
