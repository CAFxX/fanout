package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8; F1 bool; F2 uint }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int64 }

func main() {
    var v1 float32 = (0.0 * 0.0)
    v2 := make(map[uint32]int)
    var v3 int64 = -1
    var v4 int = 445
    var v5 uint64 = 2
    var v6 uint = 864
    v7 := GenS1{F0: v3}
    var v8 uint32 = 2147483648
    var v9 int32 = -100
    v10 := GenS1{F0: v3}
    v11 := GenS2{F0: -1, F1: false, F2: (^(v6))}
    for fi12 := 0; fi12 < 20; fi12++ {
        var v13 int64 = 960
        if true {
            v6 = (v6 << ((65535) & 63))
        } else {
            v1 = float32((1.5 + float64(v1)))
            fmt.Println(fi12, v4, v6)
        }
        v13 = v3
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 1; fi14++ {
        v9 -= -2147483648
        var v15 uint8 = 42
        for fi16 := 0; fi16 < 20; fi16++ {
            v5 += 42
            var v17 uint32 = (v8 & (v8 ^ v8))
            _ = fi16
            _ = v17
        }
        _ = fi14
        _ = v15
    }
    v5 -= (^(v5))
    var v18 uint16 = 42
    if true {
        for fi19 := 0; fi19 < 20; fi19++ {
            v4 = fi19
            v20 := GenS3{F0: (-(v3))}
            _ = fi19
            _ = v20
        }
        v4 = v4
        switch (("hello" + (("\n\t" + ((("\n\t" + ("hello" + ("a" + "hello"))) + "hello") + ("\n\t" + ""))) + (("a" + "foo bar") + (("" + "") + "\n\t")))) + "") {
        case "a":
            v6 -= (v6 - (v6 | (v6 ^ v6)))
        case "\n\t":
            var v21 int32 = (v9 << ((uint(v5)) & 63))
            _ = v21
        }
    }
    for fi22 := 0; fi22 < 18; fi22++ {
        if (!true) {
            for fi23 := 0; fi23 < 4; fi23++ {
                if false {
                    if true {
                        var v24 int = 32768
                        v1 = v1
                        var v25 float32 = (v1 * (0.0 + 0.0))
                        _ = v24
                        _ = v25
                    }
                } else {
                    v26 := []int{(fi22 ^ int(v18)), (^(fi23))}
                    fmt.Println(v8)
                    _ = v26
                }
                var v27 uint = v6
                if false {
                    var v28 int64 = v3
                    v27 = 0
                    if false {
                        v29 := GenS4{F0: 9223372036854775807}
                        var v30 uint16 = (v18 | 42)
                        _ = v29
                        _ = v30
                    } else {
                        v31 := []int{}
                        _ = v31
                    }
                    var v32 float32 = (float32(((0.0 + 1.5) / (-2.25 / 1.5))) - 0.0)
                    _ = v28
                    _ = v32
                } else {
                    if (v4 != (fi23 | v4)) {
                        var v33 uint32 = v8
                        var v34 float64 = 0.1
                        if (true && ((1 == v5) && (1 == v9))) {
                            var v35 float64 = v34
                            _ = v35
                        } else {
                            v36 := v2[uint32(v5)]
                            _ = v36
                        }
                        _ = v33
                        _ = v34
                    } else {
                        v37 := []int{(fi23 & (v4 * (v4 + 42))), (v4 >> ((v6) & 63))}
                        v1 = v1
                        var v38 uint8 = 255
                        _ = v37
                        _ = v38
                    }
                }
                _ = fi23
                _ = v27
            }
        } else {
            if true {
                var v39 uint32 = 2147483648
                var v40 int16 = 32767
                v3 = (v3 - (v3 << ((2) & 63)))
                _ = v39
                _ = v40
            }
            v41 := []int{}
            _ = v41
        }
        fmt.Println(v3, v18, v18)
        _ = fi22
    }
    v3 = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
