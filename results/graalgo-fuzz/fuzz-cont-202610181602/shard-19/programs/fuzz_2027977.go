package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
func genFn0(p0 uint8, p1 uint64) uint16 {
    for fi7 := 0; fi7 < 19; fi7++ {
        v8 := []int{fi7, (fi7 | -329), fi7, 128}
        if false {
            for fi9 := 0; fi9 < 12; fi9++ {
                switch fi9 {
                case -2:
                    p1 = (p1 * (p1 << ((65535) & 63)))
                default:
                    v10 := GenS1{F0: -165, F1: false}
                    _ = v10
                }
                v11 := GenS1{F0: 1, F1: (true || false)}
                _ = fi9
                _ = v11
            }
            for fi12 := 0; fi12 < 12; fi12++ {
                var v13 uint64 = p1
                p0 = p0
                _ = fi12
                _ = v13
            }
        }
        _ = fi7
        _ = v8
    }
    for fi14 := 0; fi14 < 16; fi14++ {
        p0 += (p0 >> ((42) & 63))
        p1 = p1
        _ = fi14
    }
    _ = p0
    _ = p1
    return 255
}
func genFn1(p0 bool) uint16 {
    p0 = (!p0)
    switch (("" + "a") + (((("\n\t" + "\n\t") + ((("hello" + ("" + (("foo bar" + "") + ("\n\t" + "")))) + "\n\t") + ("hello" + ("\n\t" + ((("" + ("hello" + ("hello" + ""))) + "") + ("\n\t" + ("a" + (("\n\t" + "") + "\n\t")))))))) + "\n\t") + ("a" + ""))) {
    case "foo bar":
        var v15 bool = false
        _ = v15
    }
    var v16 int32 = 42
    v17 := GenS1{F0: -59, F1: false}
    _ = p0
    _ = v16
    _ = v17
    return 65535
}

func main() {
    var v1 int32 = -1
    v2 := []uint32{42, 661}
    var v3 bool = (!false)
    v4 := []string{}
    var v5 string = ""
    var v6 int16 = 1
    var v18 int = 32767
    for fi19 := 0; fi19 < 14; fi19++ {
        for fi20 := 0; fi20 < 10; fi20++ {
            v21 := GenS1{F0: 9223372036854775807, F1: (!v3)}
            _ = fi20
            _ = v21
        }
        var v22 uint8 = 139
        for fi23 := 0; fi23 < 1; fi23++ {
            if len(v4) > 0 {
                v4[0] = (v5 + "hello")
            }
            _ = fi23
        }
        _ = fi19
        _ = v22
    }
    v18 += v18
    v3 = true
    var v24 uint64 = 0
    var v25 string = "a"
    switch v18 {
    case -578:
        v26 := GenS1{F0: 9223372036854775807, F1: (true || (-9223372036854775808 != -9223372036854775808))}
        _ = v26
    }
    v5 = v5
    fmt.Println(v25, v1)
    var v27 uint8 = 255
    if ((^(v1)) < (v1 & 758)) {
        v3 = ((-(v6)) > (v6 | (v6 / (((v6 & 1000)) | 1))))
        var v28 float32 = 0.1
        _ = len(v4)
        _ = v28
    } else {
        for fi29 := 0; fi29 < 5; fi29++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            fmt.Println(v25, fi29, v5)
            _ = len(v2)
            _ = fi29
        }
    }
    v30 := GenS1{F0: -1, F1: false}
    v5 = ""
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v24
    _ = v25
    _ = v27
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
