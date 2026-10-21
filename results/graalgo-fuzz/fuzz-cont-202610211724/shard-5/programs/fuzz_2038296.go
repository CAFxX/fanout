package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 uint16) int8 {
    var v7 uint64 = 42
    for fi8 := 0; fi8 < 12; fi8++ {
        for fi9 := 0; fi9 < 10; fi9++ {
            v10 := GenS1{F0: (p0 / ((p0) | 1))}
            switch (((("a" + ("a" + "hello")) + (("a" + "a") + ("" + "foo bar"))) + ("hello" + ("hello" + ("\n\t" + "")))) + "a") {
            case "":
                var v11 uint16 = p0
                _ = v11
            default:
                v12 := []int{892, (fi9 % ((1) | 1)), fi8, (fi9 ^ 127)}
                _ = v12
            }
            if (true && (255 <= 0)) {
                p0 = (^(p0))
            } else {
                if true {
                    var v13 string = ("hello" + ("\n\t" + "hello"))
                    if (1 < 4294967295) {
                        var v14 int64 = 1
                        _ = v14
                    }
                    fmt.Println(v13)
                    _ = v13
                }
                var v15 float64 = float64(3.14159)
                _ = v15
            }
            _ = fi9
            _ = v10
        }
        if false {
            if true {
                v7 += v7
                var v16 int32 = int32(fi8)
                _ = v16
            } else {
                var v17 int64 = 1
                var v18 float32 = (0.0 * 0.1)
                _ = v17
                _ = v18
            }
            switch fi8 {
            case 9223372036854775807:
                var v19 float32 = 3.14159
                _ = v19
            }
        } else {
            p0 = p0
            v20 := GenS1{F0: (p0 >> ((255) & 63))}
            v21 := GenS1{F0: p0}
            _ = v20
            _ = v21
        }
        fmt.Println(p0, fi8)
        _ = fi8
    }
    if (((^(p0)) >= (p0 / (((p0 ^ (p0 % (((^(p0))) | 1)))) | 1))) || false) {
        for fi22 := 0; fi22 < 7; fi22++ {
            for fi23 := 0; fi23 < 14; fi23++ {
                switch "foo bar" {
                case "a":
                    if false {
                        var v24 int = (fi22 % (((^(fi22))) | 1))
                        _ = v24
                    } else {
                        var v25 float32 = (0.1 / (float32((1e10 + (((0.1 - 1.5) - float64((float32((0.1 / -2.25)) / 1.0))) - (1e10 / 1.5)))) - 1e10))
                        if (355 <= uint64(p0)) {
                            v7 = v7
                            v7 -= v7
                        } else {
                            v26 := []int{}
                            var v27 uint16 = p0
                            if len(v26) > 0 {
                                v26[0] = fi23
                            }
                            _ = v26
                            _ = v27
                        }
                        if true {
                            var v28 uint8 = 255
                            var v29 bool = true
                            var v30 uint64 = 0
                            _ = v28
                            _ = v29
                            _ = v30
                        } else {
                            var v31 float32 = float32(0.1)
                            _ = v31
                        }
                        _ = v25
                    }
                case "foo bar":
                    v32 := GenS1{F0: (p0 | (p0 << ((537) & 63)))}
                    _ = v32
                }
                _ = fi23
            }
            fmt.Println(v7, v7, v7)
            _ = fi22
        }
    } else {
        fmt.Println(v7, p0)
        for fi33 := 0; fi33 < 2; fi33++ {
            switch (fi33 * 127) {
            case -9223372036854775808:
                switch (fi33 + fi33) {
                case 42:
                    v34 := []int{}
                    _ = v34
                case -2147483648:
                    v35 := GenS1{F0: (p0 >> ((42) & 63))}
                    _ = v35
                }
            case 1:
                for fi36 := 0; fi36 < 20; fi36++ {
                    v7 += 559
                    switch fi33 {
                    case -9223372036854775808:
                        var v37 int32 = 1
                        _ = v37
                    case -528:
                        v38 := GenS1{F0: (p0 % ((p0) | 1))}
                        _ = v38
                    }
                    _ = fi36
                }
            }
            var v39 int32 = 2147483647
            _ = fi33
            _ = v39
        }
        fmt.Println(p0)
    }
    _ = p0
    _ = v7
    return -71
}

func main() {
    var v1 int = 672
    var v2 uint16 = 32767
    v3 := make(map[uint32]string)
    var v4 uint64 = 1
    var v5 int32 = -2147483648
    var v6 float64 = (0.1 - 0.1)
    var v40 uint32 = 4294967295
    v1 -= int(v2)
    for fi41 := 0; fi41 < 18; fi41++ {
        delete(v3, v40)
        var v42 string = ("\n\t" + ((("foo bar" + ("\n\t" + "")) + "hello") + "\n\t"))
        _ = fi41
        _ = v42
    }
    v43 := GenS1{F0: (v2 % ((90) | 1))}
    var v44 int = v1
    var v45 string = "hello"
    var v46 string = v45
    v47 := GenS1{F0: (^(v2))}
    var v48 int16 = -217
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v5
    _ = v6
}
