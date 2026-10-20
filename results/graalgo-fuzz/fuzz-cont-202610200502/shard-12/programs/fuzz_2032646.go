package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 int64; F1 uint; F2 bool }
func genFn0(p0 int64, p1 bool) uint32 {
    for fi5 := 0; fi5 < 11; fi5++ {
        if (0 == 182) {
            v6 := GenS1{F0: 482, F1: p1}
            _ = v6
        }
        for fi7 := 0; fi7 < 4; fi7++ {
            p1 = (fi5 == 4294967295)
            if ((float32(float64(1e10)) - 0.0) < 1e10) {
                var v8 int16 = -150
                var v9 uint16 = 32767
                _ = v8
                _ = v9
            }
            var v10 string = "a"
            _ = fi7
            _ = v10
        }
        switch (fi5 / (((-(fi5))) | 1)) {
        case 32767:
            fmt.Println(p1)
        case 1000:
            for fi11 := 0; fi11 < 7; fi11++ {
                if p1 {
                    p1 = p1
                    v12 := []int{-2147483648}
                    var v13 int32 = 369
                    p0 = p0
                    _ = v12
                    _ = v13
                }
                _ = fi11
            }
        }
        _ = fi5
    }
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    return 2147483647
}
func genFn1(p0 bool) uint64 {
    var v14 uint = 65535
    if p0 {
        v15 := []int{}
        p0 = p0
        for fi16 := 0; fi16 < 2; fi16++ {
            if len(v15) > 0 {
                v15[0] = (fi16 >> ((v14) & 63))
            }
            if ((213 >= 247) || (!((fi16 >> ((v14) & 63)) >= (-(fi16))))) {
                if p0 {
                    for fi17 := 0; fi17 < 16; fi17++ {
                        fmt.Println(v14, p0, p0)
                        var v18 float64 = 0.0
                        _ = fi17
                        _ = v18
                    }
                }
                if p0 {
                    if p0 {
                        var v19 string = (((("" + "hello") + ("a" + "")) + ("" + ("a" + (("hello" + ("foo bar" + "a")) + ("" + ("a" + ("\n\t" + "a"))))))) + (("hello" + "a") + ""))
                        _ = v19
                    } else {
                        var v20 bool = (p0 && p0)
                        for fi21 := 0; fi21 < 4; fi21++ {
                            var v22 int64 = 9223372036854775807
                            _ = fi21
                            _ = v22
                        }
                        _ = v20
                    }
                    v14 = (v14 % ((v14) | 1))
                    v23 := GenS2{F0: 4294967295, F1: 2147483647}
                    _ = v23
                }
            }
            _ = fi16
        }
        _ = v15
    }
    _ = p0
    _ = v14
    return 42
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[string]int)
    var v3 uint = 42
    v4 := make(map[uint32]string)
    v3 -= (v3 / (((v3 * v3)) | 1))
    v24 := GenS1{F0: genFn1(true), F1: false}
    for fi25 := 0; fi25 < 10; fi25++ {
        v3 -= 255
        _ = fi25
    }
    v26 := GenS2{F0: 2147483647, F1: -1}
    for fi27 := 0; fi27 < 20; fi27++ {
        v3 = (v3 << (((v3 & 321)) & 63))
        var v28 int16 = 32767
        _ = fi27
        _ = v28
    }
    if (69 >= 31) {
        v29 := GenS1{F0: 18446744073709551615, F1: ((true && true) && (!(-1 > 1000)))}
        _ = v29
    } else {
        var v30 float32 = 1e10
        v3 = (v3 >> ((v3) & 63))
        switch 42 {
        case 42:
            v30 += (v30 - (((float32(0.0) * v30) * v30) * float32(-2.25)))
        }
        _ = v30
    }
    fmt.Println(v3)
    v3 += 312
    for fi31 := 0; fi31 < 2; fi31++ {
        switch ("\n\t" + "a") {
        case "hello":
            var v32 float32 = ((((-2.25 - (1.5 / (float32(1.5) * float32(1.5)))) - 0.0) / float32(0.1)) + (float32(-2.25) * 1.5))
            _ = v32
        }
        _ = fi31
    }
    for fi33 := 0; fi33 < 17; fi33++ {
        fmt.Println(fi33, v3)
        v34 := GenS3{F0: -634, F1: (^(v3)), F2: true}
        v35 := GenS3{F0: -467, F1: 827, F2: true}
        _ = fi33
        _ = v34
        _ = v35
    }
    fmt.Println(v3)
    v3 = (v3 << (((v3 ^ v3)) & 63))
    v3 = 65535
    v3 = v3
    fmt.Println(v3)
    var v36 int16 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v24
    _ = v26
    _ = v3
    _ = v36
    _ = v4
}
