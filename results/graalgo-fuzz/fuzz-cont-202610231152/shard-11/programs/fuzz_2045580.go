package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64; F1 int8 }
type GenS4 struct { F0 uint32; F1 uint64 }
type GenS5 struct { F0 int; F1 string; F2 uint }
func genFn0(p0 uint16, p1 int16) uint16 {
    var v6 uint16 = (^(p0))
    v6 += (v6 - v6)
    p1 += p1
    var v7 int16 = p1
    if (((1 >= -804) || false) || (4294967295 != 375)) {
        var v8 int = 395
        p1 -= (v7 >> ((1) & 63))
        v9 := GenS1{F0: 18446744073709551615, F1: -411, F2: 376}
        _ = v8
        _ = v9
    } else {
        var v10 float64 = ((3.14159 * 3.14159) - 0.0)
        v11 := GenS2{F0: v7}
        for fi12 := 0; fi12 < 18; fi12++ {
            var v13 float64 = float64((1.5 - ((1.5 - (3.14159 * (3.14159 / 1.0))) / 1.0)))
            var v14 uint = 4294967295
            if false {
                if (!true) {
                    switch 2147483647 {
                    case 32768:
                        fmt.Println(v6)
                    case -400:
                        v15 := []int{(fi12 % ((4294967295) | 1)), (fi12 - fi12), 2147483647}
                        _ = v15
                    }
                    for fi16 := 0; fi16 < 14; fi16++ {
                        v17 := GenS2{F0: (v7 + (v7 * (v7 * (p1 + (v7 >> ((v14) & 63))))))}
                        fmt.Println(v6, v13)
                        v18 := GenS3{F0: 2, F1: 0}
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                    switch fi12 {
                    case 4294967295:
                        v19 := []int{}
                        _ = v19
                    }
                }
            }
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = v10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 65535
}

func main() {
    var v1 uint8 = 2
    var v2 uint8 = (v1 ^ 195)
    v3 := []int{32767, int(v2), 42}
    v4 := []uint32{605, 506}
    var v5 string = ("foo bar" + (("foo bar" + "hello") + "foo bar"))
    v5 = ""
    var v20 uint8 = (v1 * 255)
    var v21 int32 = -2147483648
    var v22 float64 = 0.0
    v20 = v2
    v23 := GenS3{F0: 42, F1: 127}
    v4 = append(v4, 55)
    var v24 float32 = 1e10
    var v25 float32 = v24
    for fi26 := 0; fi26 < 18; fi26++ {
        for fi27 := 0; fi27 < 14; fi27++ {
            if true {
                if (!(!(genFn0(42, 1000) >= 126))) {
                    var v28 int16 = 1
                    _ = v28
                } else {
                    switch (^(fi27)) {
                    case 42:
                        var v29 int16 = -588
                        _ = v29
                    case -1:
                        v30 := GenS4{F0: 0, F1: 1}
                        _ = v30
                    case 127:
                        var v31 float64 = (v22 / (v22 + v22))
                        _ = v31
                    }
                    var v32 float32 = v24
                    var v33 uint16 = 65535
                    var v34 int8 = 1
                    var v35 uint8 = (^(v20))
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v21 -= v21
                    _ = v32
                    _ = v33
                    _ = v34
                    _ = v35
                }
            }
            _ = fi27
        }
        _ = fi26
    }
    if false {
        v24 = v24
        v2 = v20
        v36 := GenS5{F0: -344, F1: v5, F2: 123}
        _ = v36
    }
    for fi37 := 0; fi37 < 1; fi37++ {
        if ((!(!(("foo bar" + ((v5 + v5) + v5)) == "hello"))) || false) {
            var v38 uint32 = uint32(v1)
            switch fi37 {
            case 32767:
                var v39 bool = true
                _ = v39
            }
            var v40 int64 = 1
            _ = v38
            _ = v40
        }
        _ = fi37
    }
    if (v22 > ((1e10 + 1e10) / 0.1)) {
        var v41 uint16 = genFn0(0, 307)
        for fi42 := 0; fi42 < 13; fi42++ {
            switch v5 {
            case "a":
                var v43 uint16 = v41
                _ = v43
            case "hello":
                _ = len(v3)
            }
            _ = fi42
        }
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
