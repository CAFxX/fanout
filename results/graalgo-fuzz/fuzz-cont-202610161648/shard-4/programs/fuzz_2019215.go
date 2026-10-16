package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 uint) int8 {
    if (82 > 77) {
        fmt.Println(p0)
    }
    var v4 uint64 = 4294967295
    _ = p0
    _ = v4
    return 100
}
func genFn1() int16 {
    var v5 uint8 = 64
    v5 = (v5 + v5)
    _ = v5
    return 191
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]bool)
    var v3 int16 = 42
    for fi6 := 0; fi6 < 11; fi6++ {
        for fi7 := 0; fi7 < 5; fi7++ {
            if false {
                if ("\n\t" >= "a") {
                    v8 := GenS1{F0: -32768}
                    var v9 uint8 = 2
                    v9 = v9
                    _ = v8
                    _ = v9
                }
                v10 := []int{}
                v11 := v2[2147483647]
                _ = v10
                _ = v11
            } else {
                var v12 float32 = 1.5
                _ = v12
            }
            var v13 int8 = -128
            v13 = 116
            _ = fi7
            _ = v13
        }
        v14 := GenS1{F0: v3}
        var v15 float64 = 1.5
        _ = fi6
        _ = v14
        _ = v15
    }
    for fi16 := 0; fi16 < 4; fi16++ {
        switch "a" {
        case "foo bar":
            var v17 float64 = 1.5
            _ = v17
        case "hello":
            v18 := v1[(fi16 * (fi16 + fi16))]
            _ = v18
        default:
            var v19 uint8 = 251
            _ = v19
        }
        _ = fi16
    }
    if (484 < 317) {
        var v20 int16 = -32768
        _ = v20
    } else {
        if true {
            v21 := []int{}
            fmt.Println(v3)
            _ = v21
        } else {
            if true {
                var v22 int16 = v3
                for fi23 := 0; fi23 < 6; fi23++ {
                    v3 += v22
                    _ = fi23
                }
                var v24 float64 = 3.14159
                _ = v22
                _ = v24
            } else {
                if (3.14159 != float32((1e10 / 0.1))) {
                    var v25 int64 = 742
                    var v26 uint = 0
                    var v27 int16 = genFn1()
                    _ = v25
                    _ = v26
                    _ = v27
                }
                if (false && false) {
                    v3 += (-(v3))
                    var v28 uint = 255
                    for fi29 := 0; fi29 < 7; fi29++ {
                        if (true || false) {
                            v2[697] = (!(("foo bar" + (((("a" + "a") + "") + "hello") + "a")) == ("a" + "a")))
                            v3 += (v3 ^ genFn1())
                        } else {
                            v30 := GenS1{F0: (v3 & 280)}
                            v31 := GenS1{F0: genFn1()}
                            v32 := GenS1{F0: v3}
                            _ = v30
                            _ = v31
                            _ = v32
                        }
                        _ = fi29
                    }
                    _ = v28
                }
                v3 += v3
            }
        }
    }
    for fi33 := 0; fi33 < 12; fi33++ {
        fmt.Println(fi33, fi33)
        _ = fi33
    }
    var v34 uint16 = 32767
    v35 := GenS1{F0: v3}
    v36 := v2[790]
    v34 = 2
    if (3.14159 < 0.0) {
        fmt.Println(v3)
    } else {
        switch (("" + "foo bar") + ("foo bar" + ("" + ((("\n\t" + "") + (("foo bar" + ("\n\t" + ("hello" + "foo bar"))) + "a")) + "\n\t")))) {
        case "hello":
            var v37 uint32 = 0
            _ = v37
        }
        v38 := v1[-208]
        _ = v38
    }
    v39 := []int{597}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v39:", len(v39), v39)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v39
}
