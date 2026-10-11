package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 int; F2 int64 }
func genFn0(p0 int8, p1 uint16) string {
    for fi4 := 0; fi4 < 16; fi4++ {
        v5 := []int{fi4}
        if len(v5) > 0 {
            v5[0] = (^(fi4))
        }
        v6 := GenS1{F0: (p1 + (p1 * (p1 & (p1 >> ((799) & 63)))))}
        _ = fi4
        _ = v5
        _ = v6
    }
    for fi7 := 0; fi7 < 13; fi7++ {
        for fi8 := 0; fi8 < 13; fi8++ {
            switch (^(fi8)) {
            case 1:
                var v9 float64 = (3.14159 / float64((1.5 + float32(3.14159))))
                _ = v9
            case -2147483648:
                p1 = 65535
            case 166:
                v10 := GenS1{F0: 65535}
                _ = v10
            }
            var v11 int64 = int64(p1)
            var v12 float32 = float32((1e10 + -2.25))
            v13 := []int{4294967295, fi8, (fi7 ^ fi7)}
            _ = fi8
            _ = v11
            _ = v12
            _ = v13
        }
        p0 += 1
        _ = fi7
    }
    v14 := GenS1{F0: p1}
    v15 := GenS2{F0: p1, F1: 127, F2: 0}
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    return ""
}

func main() {
    v1 := make(map[uint32]string)
    var v2 float32 = 1.5
    var v3 uint = 0
    for fi16 := 0; fi16 < 8; fi16++ {
        for fi17 := 0; fi17 < 5; fi17++ {
            v18 := []int{(fi16 % ((fi17) | 1))}
            if false {
                var v19 uint8 = 67
                if (2 > (v19 >> ((v3) & 63))) {
                    if true {
                        var v20 float32 = v2
                        v2 += v2
                        _ = v20
                    }
                    fmt.Println(fi16, v19, fi16)
                    v21 := GenS2{F0: 255, F1: (fi17 >> ((65535) & 63)), F2: -9223372036854775808}
                    _ = v21
                }
                if len(v18) >= 2 {
                    v18 = v18[1:len(v18)]
                }
                _ = v19
            }
            v2 -= v2
            _ = fi17
            _ = v18
        }
        v2 = float32(-2.25)
        v22 := []int{fi16}
        _ = fi16
        _ = v22
    }
    var v23 string = ""
    for fi24 := 0; fi24 < 6; fi24++ {
        if false {
            for fi25 := 0; fi25 < 18; fi25++ {
                var v26 int8 = 100
                fmt.Println(v3, v26, v3)
                switch (fi24 >> ((v3) & 63)) {
                case 2147483647:
                    fmt.Println(v23, fi25)
                }
                _ = fi25
                _ = v26
            }
            switch v23 {
            case "hello":
                if true {
                    v27 := GenS1{F0: 2}
                    var v28 int = 385
                    var v29 uint32 = 443
                    _ = v27
                    _ = v28
                    _ = v29
                }
            case "":
                v3 = v3
            default:
                v3 -= v3
            }
            var v30 uint = (v3 & (v3 & v3))
            _ = v30
        }
        _ = fi24
    }
    switch 971 {
    case 0:
        var v31 int64 = 42
        _ = v31
    case 32768:
        v3 -= v3
    case -1:
        v3 = 801
    }
    v23 = (v23 + "\n\t")
    var v32 uint8 = 255
    v1[4294967295] = ((((v23 + (v23 + "\n\t")) + v23) + "a") + v23)
    fmt.Println(v32, v3)
    var v33 float32 = float32((1.5 + 0.1))
    v33 = v2
    var v34 uint16 = 916
    v33 = (v2 + v33)
    v33 += (0.1 - v33)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v32
    _ = v33
    _ = v34
}
