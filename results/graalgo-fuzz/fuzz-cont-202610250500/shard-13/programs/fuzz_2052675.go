package main

import (
    "fmt"
)


func main() {
    var v1 bool = (!true)
    v2 := make(map[uint32]bool)
    v3 := make(map[string]bool)
    v4 := make(map[string]bool)
    v5 := []string{}
    v1 = false
    v1 = v1
    for fi6 := 0; fi6 < 18; fi6++ {
        var v7 int = (fi6 ^ fi6)
        _ = fi6
        _ = v7
    }
    var v8 int64 = -9223372036854775808
    var v9 float32 = -2.25
    for fi10 := 0; fi10 < 3; fi10++ {
        v1 = (((true || (42 > -551)) && true) || v1)
        for fi11 := 0; fi11 < 2; fi11++ {
            switch fi10 {
            case 4294967295:
                v9 = (v9 / float32(((float64(0.1) + ((3.14159 / -2.25) - ((0.0 * float64(float32((3.14159 + (float64(float32(1e10)) / 1.0))))) + 1.5))) / 1.0)))
            }
            v9 = ((float32((1.5 * 3.14159)) / 0.1) - (v9 * ((0.1 * float32(float64((v9 * ((v9 * v9) * (v9 - 1e10)))))) + v9)))
            var v12 int32 = 166
            _ = fi11
            _ = v12
        }
        switch (^(fi10)) {
        case 42:
            switch "foo bar" {
            case "":
                v8 = -810
            case "foo bar":
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
        case 1:
            v4["\n\t"] = (-32768 > 835)
        }
        _ = fi10
    }
    var v13 bool = ((float32(-2.25) > (v9 * (v9 * v9))) && (!(v1 || (float64(1.5) >= 0.0))))
    v2[360] = (v8 >= -1)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
