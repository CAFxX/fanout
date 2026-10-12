package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 string = ((("a" + "foo bar") + ("\n\t" + "\n\t")) + "foo bar")
    var v2 bool = (!true)
    var v3 float64 = (((float64(-2.25) - 0.1) * 0.1) * ((float64(float32(0.1)) / (0.1 * 1.5)) + 3.14159))
    var v4 int64 = -1
    var v5 int32 = -475
    for fi6 := 0; fi6 < 5; fi6++ {
        v7 := []int{fi6, 844, fi6, (fi6 | fi6)}
        for fi8 := 0; fi8 < 14; fi8++ {
            v9 := GenS1{F0: 2}
            var v10 int32 = v5
            _ = fi8
            _ = v10
            _ = v9
        }
        var v11 string = v1
        _ = fi6
        _ = v11
        _ = v7
    }
    switch v1 {
    case "":
        v12 := GenS2{F0: 2147483647}
        _ = v12
    }
    var v13 uint16 = 441
    v1 = "a"
    var v14 float64 = v3
    v5 = v5
    v15 := GenS1{F0: 671}
    var v16 bool = true
    v2 = v2
    v3 -= 1e10
    fmt.Println(v14, v13)
    v16 = ("foo bar" <= v1)
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
