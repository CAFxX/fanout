package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    v1 := make(map[string]string)
    var v2 uint64 = 612
    var v3 uint32 = 42
    var v4 uint8 = 127
    if (0.0 != 1e10) {
        v3 -= v3
        for fi5 := 0; fi5 < 19; fi5++ {
            var v6 string = "foo bar"
            if true {
                delete(v1, "foo bar")
                for fi7 := 0; fi7 < 5; fi7++ {
                    fmt.Println(fi7)
                    if true {
                        v2 = 656
                    }
                    _ = fi7
                }
            }
            var v8 float32 = 0.0
            _ = fi5
            _ = v6
            _ = v8
        }
        switch "\n\t" {
        case "\n\t":
            var v9 float32 = 1.5
            _ = v9
        case "foo bar":
            v10 := GenS1{F0: v3}
            _ = v10
        case "a":
            v11 := []int{654, 4294967295}
            _ = v11
        }
    }
    var v12 uint = 0
    v13 := GenS1{F0: (^(v3))}
    for fi14 := 0; fi14 < 15; fi14++ {
        v15 := []int{fi14, (fi14 % ((fi14) | 1)), fi14, (fi14 >> ((v12) & 63))}
        _ = fi14
        _ = v15
    }
    delete(v1, "")
    var v16 int16 = int16(v4)
    var v17 int8 = 1
    if ("\n\t" < "hello") {
        v18 := []int{-1}
        fmt.Println(v17, v2, v3)
        switch 32768 {
        case 121:
            switch ("\n\t" + ((("" + "") + "\n\t") + "\n\t")) {
            case "hello":
                v17 += (v17 % ((v17) | 1))
            case "a":
                var v19 int32 = 1
                _ = v19
            }
        case 32767:
            v2 += 4294967295
        case 67:
            var v20 float64 = ((1e10 - (-2.25 * ((0.1 * -2.25) + (-2.25 / (1e10 + 1.5))))) / (-2.25 / 1e10))
            _ = v20
        default:
            if ((((v3 > 65535) && false) && true) && true) {
                v2 = (v2 * v2)
                if ((int32(v3) > 0) || (!false)) {
                    if false {
                        if true {
                            fmt.Println(v3, v4)
                        }
                    }
                    v12 -= 0
                } else {
                    v3 = v3
                }
            }
        }
        _ = v18
    } else {
        v2 = uint64(v12)
        for fi21 := 0; fi21 < 17; fi21++ {
            for fi22 := 0; fi22 < 20; fi22++ {
                v23 := []int{0, (^(fi21))}
                fmt.Println(v16, v2)
                fmt.Println(v17, v3, v3)
                if false {
                    _ = len(v1)
                    var v24 string = "\n\t"
                    _ = v24
                }
                var v25 int16 = (v16 % (((v16 & v16)) | 1))
                _ = fi22
                _ = v23
                _ = v25
            }
            v26 := GenS1{F0: uint32(v2)}
            _ = fi21
            _ = v26
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
