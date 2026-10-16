package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    v1 := []string{"hello"}
    var v2 uint = 516
    var v3 uint = (v2 * 42)
    var v4 uint16 = 32767
    v2 -= 255
    for fi5 := 0; fi5 < 6; fi5++ {
        var v6 float64 = 0.1
        v7 := GenS1{F0: -1}
        v8 := GenS1{F0: 127}
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    var v9 bool = (-1 < 1)
    if true {
        if ((-1 != 1) || false) {
            for fi10 := 0; fi10 < 20; fi10++ {
                var v11 bool = (0 >= 0)
                _ = fi10
                _ = v11
            }
            for fi12 := 0; fi12 < 12; fi12++ {
                var v13 float64 = 1e10
                v2 -= v3
                var v14 float64 = (v13 - -2.25)
                _ = fi12
                _ = v13
                _ = v14
            }
            v9 = v9
        }
        for fi15 := 0; fi15 < 6; fi15++ {
            switch 1 {
            case -2147483648:
                var v16 string = "foo bar"
                _ = v16
            case 635:
                if len(v1) > 0 {
                    v1[0] = ("a" + "\n\t")
                }
            }
            for fi17 := 0; fi17 < 10; fi17++ {
                v18 := GenS1{F0: 100}
                _ = fi17
                _ = v18
            }
            _ = fi15
        }
    } else {
        v9 = ((2 > 42) || v9)
        var v19 float64 = 3.14159
        v9 = (554 >= 42)
        _ = v19
    }
    var v20 uint8 = 128
    var v21 uint16 = v4
    var v22 uint8 = 1
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v9
}
