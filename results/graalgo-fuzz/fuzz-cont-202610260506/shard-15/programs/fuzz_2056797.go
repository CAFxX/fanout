package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 int8 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint = 4294967295
    v3 := make(map[uint32]uint32)
    v4 := []string{("foo bar" + "hello"), ("a" + "hello"), "foo bar", (("a" + "\n\t") + ("\n\t" + (("" + "hello") + "foo bar")))}
    v5 := GenS1{F0: "a", F1: v2, F2: 100}
    for fi6 := 0; fi6 < 4; fi6++ {
        for fi7 := 0; fi7 < 14; fi7++ {
            var v8 float32 = (1e10 + 1e10)
            _ = fi7
            _ = v8
        }
        v9 := GenS1{F0: ("\n\t" + "\n\t"), F1: (v2 | v2), F2: 127}
        v4 = append(v4, "\n\t")
        _ = fi6
        _ = v9
    }
    v4 = append(v4, "")
    for fi10 := 0; fi10 < 11; fi10++ {
        delete(v3, 2147483647)
        var v11 bool = ((42 <= 1) && (181 < 1))
        for fi12 := 0; fi12 < 19; fi12++ {
            v11 = false
            _ = fi12
        }
        _ = fi10
        _ = v11
    }
    var v13 string = ("hello" + (("a" + "hello") + ""))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v14 int32 = -2147483648
    if false {
        delete(v1, 2)
        fmt.Println(v14)
        fmt.Println(v13)
    } else {
        var v15 float64 = -2.25
        if false {
            switch -9223372036854775808 {
            case 2147483647:
                _ = len(v4)
            case 167:
                if (true && true) {
                    v4 = append(v4, v13)
                }
            }
            for fi16 := 0; fi16 < 18; fi16++ {
                v4 = append(v4, v13)
                _ = fi16
            }
        }
        if len(v4) > 0 {
            v4[0] = "foo bar"
        }
        _ = v15
    }
    var v17 uint16 = 512
    delete(v1, 2147483647)
    for fi18 := 0; fi18 < 13; fi18++ {
        v19 := GenS1{F0: "a", F1: (v2 & 246), F2: 1}
        var v20 float32 = (1.5 - ((float32((1.5 / 3.14159)) / ((-2.25 + 0.1) + float32(((1e10 / float64(1.5)) / 1e10)))) * ((((float32(3.14159) - 3.14159) / ((0.0 - 3.14159) - float32(1e10))) / -2.25) + (3.14159 + (0.0 / 1.0)))))
        var v21 int16 = -428
        _ = fi18
        _ = v19
        _ = v20
        _ = v21
    }
    var v22 int8 = 127
    for fi23 := 0; fi23 < 20; fi23++ {
        for fi24 := 0; fi24 < 2; fi24++ {
            v22 += v22
            switch fi24 {
            case 2147483647:
                fmt.Println(v13, v2)
            }
            switch v13 {
            case "\n\t":
                v25 := GenS1{F0: (v13 + "foo bar"), F1: (v2 >> ((v2) & 63)), F2: (v22 / ((v22) | 1))}
                _ = v25
            case "foo bar":
                var v26 int32 = -404
                _ = v26
            default:
                var v27 int = (-(fi23))
                _ = v27
            }
            _ = len(v4)
            var v28 uint8 = 176
            _ = fi24
            _ = v28
        }
        var v29 int8 = (v22 / ((0) | 1))
        _ = fi23
        _ = v29
    }
    v4 = append(v4, (v13 + "foo bar"))
    var v30 int16 = -1
    var v31 uint64 = 1
    v30 -= (v30 + 0)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v22
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
