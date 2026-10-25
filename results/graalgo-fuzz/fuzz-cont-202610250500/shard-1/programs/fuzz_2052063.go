package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64 }
func genFn0() int32 {
    switch (((("" + "\n\t") + "\n\t") + "foo bar") + (("\n\t" + "foo bar") + ("\n\t" + "foo bar"))) {
    case "\n\t":
        v5 := GenS1{F0: -386}
        _ = v5
    }
    var v6 float32 = 3.14159
    var v7 uint32 = 1
    v6 += ((v6 * ((v6 - 0.0) * ((-2.25 - v6) * ((v6 / 1e10) - 0.0)))) - v6)
    _ = v6
    _ = v7
    return 0
}

func main() {
    var v1 uint32 = 2147483648
    var v2 bool = (v1 != (v1 & (v1 * (v1 | (v1 | (v1 % (((v1 - (v1 * (v1 >> ((255) & 63))))) | 1)))))))
    var v3 int = -2147483648
    v4 := make(map[string]int)
    if (0 <= 0) {
        var v8 string = "a"
        _ = v8
    } else {
        var v9 uint64 = 42
        for fi10 := 0; fi10 < 10; fi10++ {
            var v11 int64 = 42
            v12 := GenS1{F0: -579}
            var v13 int8 = 1
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
        }
        _ = v9
    }
    for fi14 := 0; fi14 < 1; fi14++ {
        v3 = fi14
        if false {
            if (!(!v2)) {
                if (-546 != 0) {
                    switch (((("" + "a") + "a") + "\n\t") + ((("foo bar" + "") + "\n\t") + "hello")) {
                    case "foo bar":
                        if v2 {
                            v15 := []int{(v3 * (v3 | v3)), (fi14 & (v3 >> ((0) & 63))), 1}
                            v3 -= v3
                            _ = v15
                        } else {
                            v16 := GenS2{F0: 1}
                            v2 = (v2 && (!(859 > 1)))
                            v17 := GenS2{F0: 634}
                            _ = v16
                            _ = v17
                        }
                    }
                }
                for fi18 := 0; fi18 < 9; fi18++ {
                    fmt.Println(v1, v2, v2)
                    v19 := []int{fi18, 820}
                    _ = fi18
                    _ = v19
                }
            } else {
                fmt.Println(fi14, v2, fi14)
            }
        } else {
            v20 := []int{}
            var v21 bool = false
            if (v1 <= v1) {
                if len(v20) >= 2 {
                    v20 = v20[1:len(v20)]
                }
            }
            _ = v20
            _ = v21
        }
        _ = fi14
    }
    var v22 bool = (-734 != -190)
    v3 = 412
    fmt.Println(v1, v1)
    switch (v3 >> ((234) & 63)) {
    case 336:
        for fi23 := 0; fi23 < 19; fi23++ {
            for fi24 := 0; fi24 < 18; fi24++ {
                var v25 float64 = (0.0 / -2.25)
                var v26 uint16 = 0
                v27 := v4[""]
                _ = fi24
                _ = v25
                _ = v26
                _ = v27
            }
            v28 := GenS2{F0: 228}
            _ = fi23
            _ = v28
        }
    case 32767:
        switch "\n\t" {
        case "":
            switch ("" + "foo bar") {
            case "":
                var v29 string = "a"
                _ = v29
            case "foo bar":
                var v30 uint32 = (^(v1))
                _ = v30
            default:
                var v31 int = -1
                _ = v31
            }
        }
    case 779:
        var v32 uint16 = 65535
        _ = v32
    default:
        var v33 int64 = 888
        _ = v33
    }
    var v34 int16 = -766
    var v35 uint8 = 0
    for fi36 := 0; fi36 < 15; fi36++ {
        v3 -= fi36
        _ = fi36
    }
    var v37 string = "\n\t"
    v38 := []int{32767}
    v38 = append(v38, v3)
    v39 := GenS1{F0: -2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v34
    _ = v35
    _ = v37
    _ = v38
    _ = v39
    _ = v4
}
