package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8 }

func main() {
    v1 := []string{((("\n\t" + ("a" + "\n\t")) + ("foo bar" + "")) + (("a" + "a") + "a"))}
    v2 := make(map[uint32]string)
    var v3 float32 = (1e10 / float32(-2.25))
    for fi4 := 0; fi4 < 7; fi4++ {
        if (!true) {
            fmt.Println(fi4, fi4)
            v5 := GenS1{F0: 0}
            _ = v5
        } else {
            v3 = ((1e10 + v3) * (v3 / 1e10))
            for fi6 := 0; fi6 < 4; fi6++ {
                if ((((((float64(v3) - (float64(v3) + ((-2.25 / float64(v3)) - float64((((v3 / -2.25) * (v3 * (v3 + (float32(0.1) * float32((float64(v3) - 3.14159)))))) * v3))))) + 0.1) >= (1.5 * 3.14159)) && ((v3 + v3) == (-2.25 * 1e10))) || (!false)) || false) {
                    if false {
                        v7 := v2[526]
                        v1 = append(v1, v7)
                        fmt.Println(fi4, fi4)
                        _ = v7
                    }
                    var v8 int64 = 9223372036854775807
                    _ = v8
                } else {
                    v3 = float32(0.0)
                }
                v3 -= 0.1
                _ = fi6
            }
        }
        _ = fi4
    }
    switch -9223372036854775808 {
    case 42:
        v9 := GenS1{F0: 0}
        _ = v9
    case -1:
        for fi10 := 0; fi10 < 4; fi10++ {
            switch (fi10 << ((4294967295) & 63)) {
            case -9223372036854775808:
                v3 = v3
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v11 float32 = (v3 - v3)
            _ = fi10
            _ = v11
        }
    }
    v3 -= 1e10
    v3 = (v3 - v3)
    for fi12 := 0; fi12 < 13; fi12++ {
        var v13 int8 = -1
        _ = fi12
        _ = v13
    }
    delete(v2, 1)
    var v14 int = -596
    for fi15 := 0; fi15 < 8; fi15++ {
        if true {
            for fi16 := 0; fi16 < 14; fi16++ {
                v17 := v2[2147483647]
                delete(v2, 2)
                var v18 float64 = 0.1
                if len(v1) > 0 {
                    v1[0] = v17
                }
                switch v17 {
                case "":
                    v3 = (1.5 / 1.0)
                }
                _ = fi16
                _ = v17
                _ = v18
            }
            v14 = (fi15 % (((-(fi15))) | 1))
        } else {
            v3 += (3.14159 + v3)
        }
        _ = fi15
    }
    var v19 int32 = 42
    v20 := GenS2{F0: 1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
