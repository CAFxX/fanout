package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 uint32 }
type GenS3 struct { F0 bool; F1 uint32; F2 int64 }
type GenS4 struct { F0 uint8; F1 uint16 }
func genFn0(p0 int) uint {
    var v7 float32 = -2.25
    v8 := []int{(p0 ^ 631), p0}
    v9 := GenS1{F0: 560}
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 42
}
func genFn1(p0 int64) string {
    var v10 float64 = 1.5
    v11 := GenS1{F0: -2147483648}
    fmt.Println(p0)
    _ = p0
    _ = v10
    _ = v11
    return "\n\t"
}

func main() {
    var v1 int8 = 100
    v2 := make(map[string]int)
    var v3 bool = (true || (true && true))
    var v4 int32 = 2147483647
    var v5 uint = 907
    var v6 int8 = 100
    v6 = -128
    var v12 uint16 = 255
    if v3 {
        v12 = 2
        v13 := GenS1{F0: -1}
        _ = v13
    }
    for fi14 := 0; fi14 < 19; fi14++ {
        switch (("" + ("hello" + "hello")) + "\n\t") {
        case "":
            if (!true) {
                for fi15 := 0; fi15 < 10; fi15++ {
                    v16 := []int{fi14, fi14}
                    delete(v2, ("" + "\n\t"))
                    _ = fi15
                    _ = v16
                }
            } else {
                v6 = v6
                fmt.Println(v1, v6, v1)
                for fi17 := 0; fi17 < 7; fi17++ {
                    var v18 int16 = 655
                    if (v3 && v3) {
                        var v19 int32 = (v4 % ((v4) | 1))
                        v6 -= (v6 << (((v5 & 891)) & 63))
                        _ = v19
                    }
                    var v20 uint64 = 42
                    _ = fi17
                    _ = v18
                    _ = v20
                }
            }
        case "hello":
            var v21 uint8 = 255
            _ = v21
        case "foo bar":
            if (18446744073709551615 >= 60) {
                v6 -= (v1 >> ((v5) & 63))
            }
        default:
            for fi22 := 0; fi22 < 9; fi22++ {
                if v3 {
                    v23 := v2[("\n\t" + ("" + "\n\t"))]
                    _ = v23
                }
                if true {
                    var v24 uint16 = v12
                    _ = v24
                }
                _ = fi22
            }
        }
        _ = fi14
    }
    var v25 uint = (v5 >> ((v5) & 63))
    var v26 uint64 = 42
    var v27 float64 = 0.0
    for fi28 := 0; fi28 < 3; fi28++ {
        v29 := GenS1{F0: -56}
        if v3 {
            v30 := GenS2{F0: v1, F1: 2147483647}
            v31 := GenS3{F0: (!v3), F1: 65535, F2: 800}
            var v32 uint64 = 264
            _ = v30
            _ = v31
            _ = v32
        } else {
            v33 := v2["\n\t"]
            v27 += v27
            _ = v33
        }
        switch fi28 {
        case -80:
            v4 = (v4 >> ((v25) & 63))
        case 4294967295:
            var v34 string = ""
            _ = v34
        default:
            switch (fi28 * (fi28 ^ (-(fi28)))) {
            case 833:
                if ("a" <= (("" + ((("\n\t" + "hello") + "hello") + "\n\t")) + ("hello" + ("\n\t" + "a")))) {
                    fmt.Println(v5, fi28, v5)
                    v35 := GenS2{F0: 1, F1: 42}
                    v36 := v2["\n\t"]
                    _ = v35
                    _ = v36
                }
            case -1:
                if (127 == (^(v6))) {
                    v37 := GenS4{F0: 7, F1: 42}
                    var v38 int32 = v4
                    v6 -= v6
                    _ = v37
                    _ = v38
                }
            }
        }
        _ = fi28
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
