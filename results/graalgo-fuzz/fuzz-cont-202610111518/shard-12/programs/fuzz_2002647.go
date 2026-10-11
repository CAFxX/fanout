package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 int32 }
func genFn0(p0 string) uint {
    var v5 uint32 = 4294967295
    v6 := []int{127, 32767}
    _ = p0
    _ = v5
    _ = v6
    return 0
}

func main() {
    var v1 int64 = 1
    v2 := make(map[uint32]string)
    v3 := []int{25, -512, 42, -1}
    v4 := make(map[string]int)
    for fi7 := 0; fi7 < 17; fi7++ {
        for fi8 := 0; fi8 < 4; fi8++ {
            switch "foo bar" {
            case "foo bar":
                if (!true) {
                    v9 := GenS1{F0: 0, F1: 0, F2: -2147483648}
                    _ = v9
                }
            case "a":
                var v10 float64 = 0.1
                _ = v10
            case "hello":
                _ = len(v3)
            }
            var v11 string = "foo bar"
            _ = fi8
            _ = v11
        }
        var v12 uint32 = 42
        _ = fi7
        _ = v12
    }
    if true {
        var v13 int = 254
        switch (v13 | v13) {
        case 2147483647:
            switch v13 {
            case 741:
                v13 = 4294967295
            case 127:
                v1 = (^(v1))
            case -462:
                var v14 string = "hello"
                _ = v14
            }
        }
        _ = v13
    }
    v1 = (v1 | v1)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v15 string = ("hello" + ("foo bar" + "foo bar"))
    var v16 int16 = 42
    v17 := v4[(("" + (("\n\t" + v15) + v15)) + v15)]
    if true {
        var v18 int32 = 970
        v15 = ("hello" + v15)
        v1 += (v1 / ((1) | 1))
        _ = v18
    } else {
        v19 := GenS1{F0: genFn0(v15), F1: 42, F2: 1}
        _ = v19
    }
    v17 = (v17 % ((-117) | 1))
    var v20 float32 = 0.1
    for fi21 := 0; fi21 < 10; fi21++ {
        switch (v17 / ((fi21) | 1)) {
        case 394:
            v22 := v2[2147483647]
            _ = v22
        case 32768:
            v1 = v1
        case -675:
            var v23 uint = 255
            _ = v23
        }
        if true {
            v16 = (^(v16))
            v17 += (v17 | v17)
        } else {
            for fi24 := 0; fi24 < 10; fi24++ {
                if false {
                    v1 += (v1 & v1)
                    var v25 int32 = -2147483648
                    fmt.Println(v1)
                    _ = v25
                }
                v26 := GenS1{F0: 519, F1: 65535, F2: -819}
                v20 = v20
                _ = fi24
                _ = v26
            }
            var v27 int32 = 42
            v28 := GenS1{F0: 494, F1: 4294967295, F2: v27}
            _ = v27
            _ = v28
        }
        _ = fi21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
