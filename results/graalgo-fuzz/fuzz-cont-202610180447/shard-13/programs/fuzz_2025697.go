package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16 }
func genFn0() string {
    if ((0 != 781) && false) {
        var v6 uint64 = 0
        v7 := GenS1{F0: (2147483647 != 2), F1: 42}
        _ = v6
        _ = v7
    }
    var v8 float64 = 0.1
    var v9 uint = 0
    fmt.Println(v8, v9)
    _ = v8
    _ = v9
    return ((("a" + "a") + "foo bar") + "hello")
}
func genFn1(p0 string, p1 uint32) int {
    v10 := GenS2{F0: -65}
    var v11 uint = 255
    for fi12 := 0; fi12 < 17; fi12++ {
        if true {
            switch (p0 + (p0 + (("\n\t" + p0) + (p0 + p0)))) {
            case "hello":
                var v13 float32 = float32(0.1)
                _ = v13
            }
        } else {
            var v14 string = "hello"
            if (fi12 == (fi12 * (fi12 >> (((v11 - v11)) & 63)))) {
                v15 := []int{fi12, 32768, (fi12 + -2147483648), fi12}
                _ = v15
            }
            _ = v14
        }
        var v16 int = fi12
        var v17 string = p0
        _ = fi12
        _ = v16
        _ = v17
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return 0
}

func main() {
    v1 := make(map[int]bool)
    v2 := []int{2147483647, 307}
    var v3 int32 = -2147483648
    var v4 uint16 = 1
    var v5 int = 399
    var v18 int8 = 42
    delete(v1, int(v3))
    var v19 float32 = 1e10
    v5 = (v5 - (v5 << ((742) & 63)))
    fmt.Println(v4)
    fmt.Println(v19)
    for fi20 := 0; fi20 < 18; fi20++ {
        v19 = -2.25
        _ = fi20
    }
    v21 := GenS1{F0: ((float32(((((0.1 + 0.1) - 3.14159) + float64(1.5)) / float64(((1.5 / (3.14159 / (3.14159 - v19))) * (-2.25 / v19))))) >= 1e10) && false), F1: v4}
    var v22 uint32 = 65535
    var v23 int8 = (-(v18))
    for fi24 := 0; fi24 < 1; fi24++ {
        for fi25 := 0; fi25 < 4; fi25++ {
            if (v4 == (v4 % (((v4 * 327)) | 1))) {
                v19 -= 0.0
                v23 = (v23 >> ((255) & 63))
                v26 := GenS1{F0: true, F1: (^(v4))}
                fmt.Println(v5)
                v27 := GenS2{F0: (^(v23))}
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = v26
                _ = v27
            } else {
                if false {
                    var v28 int = fi24
                    _ = v28
                }
                v18 = 100
                _ = len(v1)
            }
            var v29 bool = (!(!((!false) && true)))
            if (1000 == -1) {
                v18 = 1
                v29 = v29
                var v30 float32 = 1e10
                switch (v5 << ((18446744073709551615) & 63)) {
                case 0:
                    var v31 int32 = -668
                    _ = v31
                case -468:
                    var v32 uint8 = 196
                    _ = v32
                }
                v33 := GenS3{F0: v4}
                var v34 uint = 65535
                _ = v30
                _ = v33
                _ = v34
            } else {
                var v35 int8 = (-(v18))
                switch fi25 {
                case -2147483648:
                    v5 = v5
                }
                var v36 string = (("" + "a") + "\n\t")
                _ = v35
                _ = v36
            }
            _ = fi25
            _ = v29
        }
        if (true && (!(true && ((true && false) && true)))) {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        v1[fi24] = true
        _ = fi24
    }
    if (!false) {
        var v37 int = (v5 % ((v5) | 1))
        v37 = 32767
        _ = v37
    } else {
        _ = len(v2)
        var v38 int8 = (-(v23))
        _ = v38
    }
    var v39 int8 = v18
    if true {
        switch "" {
        case "hello":
            var v40 float64 = -2.25
            _ = v40
        }
        switch genFn1("hello", (v22 >> ((65535) & 63))) {
        case 955:
            var v41 uint64 = 722
            _ = v41
        }
    } else {
        fmt.Println(v39, v22)
        v19 += (float32(-2.25) * (((((0.1 + 0.1) + 1.5) + (v19 * ((v19 - v19) / float32((1.5 * float64(v19)))))) * v19) + ((((v19 + (v19 + (v19 / -2.25))) * (0.0 * ((-2.25 - v19) + v19))) * 0.0) - ((v19 + v19) - v19))))
        v42 := GenS2{F0: v39}
        _ = v42
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v39
    _ = v4
    _ = v5
}
