package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 int32 }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 uint8; F1 string; F2 uint8 }
func genFn0() int64 {
    for fi4 := 0; fi4 < 18; fi4++ {
        if true {
            var v5 int = fi4
            for fi6 := 0; fi6 < 7; fi6++ {
                v7 := GenS1{F0: 1, F1: 1, F2: 2147483647}
                var v8 int = 2147483647
                v9 := []int{(-(fi4)), (^(fi4)), (-(fi6))}
                v8 -= (fi4 ^ (fi4 << ((65535) & 63)))
                var v10 uint32 = 749
                _ = fi6
                _ = v10
                _ = v7
                _ = v8
                _ = v9
            }
            switch (fi4 & v5) {
            case 128:
                var v11 uint16 = 255
                _ = v11
            case 127:
                v5 = fi4
            case 9223372036854775807:
                if ((1 <= 0) && false) {
                    for fi12 := 0; fi12 < 6; fi12++ {
                        v13 := GenS1{F0: 2147483647, F1: 1, F2: -667}
                        _ = fi12
                        _ = v13
                    }
                    switch v5 {
                    case 4294967295:
                        if false {
                            var v14 uint8 = 128
                            _ = v14
                        }
                    case -183:
                        for fi15 := 0; fi15 < 7; fi15++ {
                            v16 := []int{}
                            v17 := GenS2{F0: (-(fi4)), F1: 23}
                            _ = fi15
                            _ = v16
                            _ = v17
                        }
                    case -1:
                        if (1 == 841) {
                            fmt.Println(fi4)
                        }
                    }
                    v18 := []int{(v5 - v5)}
                    _ = v18
                }
            default:
                for fi19 := 0; fi19 < 3; fi19++ {
                    if false {
                        if true {
                            fmt.Println(v5, v5)
                        }
                    }
                    v20 := []int{(fi4 - fi4), (v5 + (fi19 % ((fi4) | 1))), (v5 | (v5 >> ((65535) & 63)))}
                    _ = fi19
                    _ = v20
                }
            }
            _ = v5
        }
        _ = fi4
    }
    var v21 float64 = float64((1.5 + -2.25))
    v21 = v21
    v22 := GenS1{F0: -201, F1: 2, F2: -1}
    var v23 uint16 = 42
    _ = v21
    _ = v22
    _ = v23
    return 414
}

func main() {
    v1 := make(map[string]uint32)
    var v2 float64 = (0.1 + 0.0)
    var v3 uint32 = 284
    fmt.Println(v3, v2)
    if (0.0 > (((float32((float64(1.5) + 0.0)) - ((3.14159 + 1e10) * 1.5)) * float32(v2)) + 1e10)) {
        if (true && false) {
            for fi24 := 0; fi24 < 19; fi24++ {
                v25 := []int{fi24}
                _ = fi24
                _ = v25
            }
            var v26 uint = 129
            _ = v26
        }
        var v27 uint = 558
        _ = v27
    } else {
        v2 = ((v2 * float64((3.14159 / ((-2.25 / 3.14159) + (0.0 / 1e10))))) + float64(1e10))
        switch ("hello" + "hello") {
        case "":
            for fi28 := 0; fi28 < 2; fi28++ {
                if true {
                    v29 := GenS1{F0: 1, F1: 741, F2: -494}
                    v3 -= 4294967295
                    var v30 uint64 = 2
                    _ = v29
                    _ = v30
                }
                _ = fi28
            }
        }
    }
    v2 = v2
    v1["foo bar"] = v3
    for fi31 := 0; fi31 < 18; fi31++ {
        fmt.Println(fi31, fi31, v3)
        var v32 int64 = genFn0()
        _ = fi31
        _ = v32
    }
    if false {
        if (-317 <= 2147483647) {
            var v33 uint = 330
            _ = v33
        }
    }
    var v34 string = "a"
    var v35 int16 = 32767
    v3 = v3
    if (!(true || (!true))) {
        var v36 float64 = 3.14159
        v35 = v35
        _ = v36
    } else {
        fmt.Println(v34)
        v1[(v34 + "\n\t")] = 948
    }
    var v37 float64 = v2
    fmt.Println(v35, v37)
    fmt.Println(v37, v34)
    var v38 string = ""
    if (v38 == v34) {
        _ = len(v1)
        v39 := v1[v38]
        _ = v39
    } else {
        switch v34 {
        case "foo bar":
            v35 -= v35
        case "":
            v40 := GenS3{F0: 127, F1: "hello", F2: 175}
            _ = v40
        default:
            for fi41 := 0; fi41 < 6; fi41++ {
                v42 := []int{}
                _ = fi41
                _ = v42
            }
        }
        var v43 uint64 = 1
        _ = v43
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v37
    _ = v38
}
